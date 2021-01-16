package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"novel/app/cache"
	"novel/app/cache/cache_key"
	"novel/app/dao"
	"novel/app/dto"
	"novel/app/global/consts"
	"novel/app/log"
	"novel/app/model"
)

var bookDao = dao.Book{}
var bookIndexDao = dao.BookIndex{}
var bookContentDao = dao.BookContent{}

type BookService struct {
}

func (*BookService) Create(dto dto.BookCreateDto) (*model.Book, error) {

	// 先查询, 书名不能重复
	if bookDao.GetByBookName(dto.BookName).BookName == dto.BookName {
		return nil, errors.New("书名不能重复")
	}

	bookModel := model.Book{}
	// 这里使用copy模块直接简化属性复制的操作, 利用的反射复制的性能代价非常小
	copier.Copy(&bookModel, &dto)

	c := bookDao.Create(&bookModel)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return &bookModel, c.Error
	}
	return &bookModel, nil
}

func (*BookService) GetBookDetail(id int64) (*model.Book, error) {
	book := bookDao.GetByBookId(id)
	if book.Id <= 0 {
		return nil, errors.New("不存在当前书籍")
	}
	return &book, nil
}

func (*BookService) List() ([]model.Book, int64) {
	books, total := bookDao.List()
	return books, total
}

func (*BookService) ListClickRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexClickBankBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(0, 10)
		_ = cache.SetStruct(cache_key.IndexClickBankBookKey, &books, cache_key.IndexBookKeyTime)
	}
	return books
}

func (*BookService) ListRank(rankType int8, limit int) []model.Book {
	return bookDao.ListRank(rankType, limit)
}

func (*BookService) ListNewRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexNewBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(1, 10)
		err := cache.SetStruct(cache_key.IndexNewBookKey, &books, cache_key.IndexBookKeyTime)
		fmt.Println(err, "err")
	}
	return books
}

func (*BookService) ListUpdateRank() []model.Book {
	var books []model.Book
	_ = cache.GetStruct(cache_key.IndexUpdateBookKey, &books)
	if books == nil {
		books = bookDao.ListRank(2, 23)
		_ = cache.SetStruct(cache_key.IndexUpdateBookKey, &books, cache_key.IndexBookKeyTime)
	}
	return books
}

func (*BookService) ListBookCategory() []dto.BookCategoryRespDto {
	return bookDao.ListBookCategory()
}

func (*BookService) AddVisitCount(bookId int64, visitCount int) {
	bookDao.AddVisitCount(bookId, visitCount)
}

func (*BookService) QueryIndexCount(bookId int64) int64 {
	return bookIndexDao.QueryIndexCount(bookId)
}

func (*BookService) QueryBookContent(bookIndexId int64) (*model.BookContent, error) {
	bookContent := bookContentDao.QueryBookContent(bookIndexId)
	if bookContent.ID <= 0 {
		return nil, errors.New("章节不存在")
	}
	return bookContent, nil
}

func (bs *BookService) QueryBookIndexAbout(bookId int64, bookIndexId int64, isCut bool) (*dto.BookIndexAboutRespDto, error) {
	bookContent, err := bs.QueryBookContent(bookIndexId)
	if err != nil {
		return nil, err
	}
	count := bookIndexDao.QueryIndexCount(bookId)
	content := bookContent.Content
	if isCut {
		cutLen := 120 // 裁剪的长度
		if len(content) > cutLen {
			content = content[0:cutLen]
		}
	}
	// 查询书籍首个章节id信息
	firstId := bookIndexDao.GetFirstIndexId(bookId)
	return &dto.BookIndexAboutRespDto{
		BookIndexCount:   count,
		BookContent:      content,
		BookFirstIndexId: firstId,
	}, nil
}

func (bs *BookService) ListRecBookByCatId(bookId int64, catId int64) []dto.ListRecBookRespDto {
	// todo: 根据分类id查询, 暂无推荐算法, 随机推荐同类书籍
	return bookDao.ListRecBookByCatId(bookId, catId)
}

func (bs *BookService) ListCommentByPage(userId int64, bookId int64, page int, pageSize int) ([]dto.ListCommentRespDto, int64) {
	res, count := bookDao.ListCommentByPage(userId, bookId, page, pageSize)
	return res, count
}

func (bs *BookService) AddBookComment(ccDto dto.CommentCreateDto, userId int64) error {
	// 检查用户是否有评价过书籍
	comment := bookDao.GetUserComment(ccDto.BookId, userId)
	var modelComment model.BookComment
	var err error
	var tx *gorm.DB
	defer func() {
		if err != nil && tx != nil {
			// 统一回滚事务
			tx.Rollback()
		}
	}()

	if comment.ID > 0 {
		return errors.New("已经评价过该书了")
	}
	// 检查书籍是否存在
	if bookDao.GetByBookId(ccDto.BookId).Id <= 0 {
		return errors.New("书籍不存在")
	}

	copier.Copy(&modelComment, &ccDto)
	modelComment.CreateUserID = userId

	// 开启事务, gorm没有函数式的事务开启方式, 只能使用这种取巧的方式了.
	tx = dao.GetDb().Begin()
	bookDao := dao.Book{DB: tx}
	if err = bookDao.AddBookComment(modelComment); err != nil {
		return errors.New("添加评价失败")
	}
	// 下面测试事务问题
	// time.Sleep(time.Second * 15)
	// return errors.New("测试事务的错误....")

	// 增加书籍评论数量
	if err = bookDao.AddCommentCount(ccDto.BookId); err != nil {
		return errors.New("添加评价失败2")
	}
	tx.Commit()
	return nil
}

// 根据书籍id获取目录列表
func (bs *BookService) QueryIndexList(bookId int64, orderBy string, page int, pageSize int) ([]dto.IndexListRespDto, int64) {
	// 这里根据orderBy的属性来选取排序条件
	var order string
	switch orderBy {
	case "-index_num":
		order = "-index_num"
	default:
		order = "index_num"
	}
	return bookDao.QueryIndexList(bookId, order, page, pageSize)
}

// 获取首页配置
func (bs *BookService) ListBookSetting() map[int8][]dto.BookSettingDto {
	settings := bookIndexDao.GetIndexSettings()
	if len(settings) == 0 {
		settings = bs.initIndexSettings()
	}
	// 将数组分类, 组成map格式
	settingsToMap := bs.listSettingsToMap(settings)
	return settingsToMap
}

func (bs *BookService) initIndexSettings() []dto.BookSettingDto {
	books := bookDao.GetBooksByScoreRandom(consts.IndexBookSettingNum)
	var settingsDto []dto.BookSettingDto
	var settingsModel []model.BookSetting
	length := len(books)
	var _sort int8
	for i := 0; i < length; i++ {
		var _type int8
		if i < 4 {
			_type = 0
			_sort = int8(i + 1)
		} else if i < 14 {
			_type = 1
			_sort = int8(i + 1 - 4)
		} else if i < 20 {
			_type = 2
			_sort = int8(i + 1 - 14)
		} else if i < 26 {
			_type = 3
			_sort = int8(i + 1 - 20)
		} else {
			_type = 4
			_sort = int8(i + 1 - 26)
		}
		var settingDto = dto.BookSettingDto{}
		var settingModel = model.BookSetting{
			BookId: books[i].Id,
			Sort:   _sort,
			Type:   _type,
		}
		copier.Copy(&settingDto, &books[i])
		settingDto.BookId = settingModel.BookId
		settingDto.Sort = settingModel.Sort
		settingDto.Type = settingModel.Type

		settingsDto = append(settingsDto, settingDto)
		settingsModel = append(settingsModel, settingModel)
	}
	// 保存首页设置
	_ = bookDao.InsertManySettings(&settingsModel)
	return settingsDto
}

func (bs *BookService) listSettingsToMap(settings []dto.BookSettingDto) map[int8][]dto.BookSettingDto {
	var settingsMap = make(map[int8][]dto.BookSettingDto)
	for _, setting := range settings {
		type_ := setting.Type
		if v, ok := settingsMap[type_]; ok {
			v = append(v, setting)
			settingsMap[type_] = v
		} else {
			settingsMap[type_] = []dto.BookSettingDto{setting}
		}
	}
	return settingsMap
}

func (bs *BookService) QueryPreBookIndexId(bookId int64, indexId int64) int64 {
	return bookDao.QueryNearBookIndexId(bookId, indexId, false)
}

func (bs *BookService) QueryNextBookIndexId(bookId int64, indexId int64) int64 {
	return bookDao.QueryNearBookIndexId(bookId, indexId, true)
}

func (bs *BookService) GetIndexDetail(indexId int64) (*model.BookIndex, error) {
	index := bookIndexDao.GetIndexDetail(indexId)
	if index.Id <= 0 {
		return nil, errors.New("不存在当前章节")
	}
	return &index, nil
}
