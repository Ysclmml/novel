// 登录
import instance from "@/utils/request";

export function listBookSetting (params) {
    return instance({
        url: '/v1/book/listBookSetting',
        method: 'get',
        params
    })
}


export function listClickRank (params) {
    return instance({
        url: '/v1/book/listClickRank',
        method: 'get',
        params
    })
}

export function listNewRank (params) {
    return instance({
        url: '/v1/book/listNewRank',
        method: 'get',
        params
    })
}

export function listUpdateRank (params) {
    return instance({
        url: '/v1/book/listUpdateRank',
        method: 'get',
        params
    })
}

// 查看排行帮更多信息, 暂时默认30条
export function listBooksRank (type, limit = 30) {
    return instance({
        url: '/v1/book/listRank',
        method: 'get',
        params: {
            type,
            limit
        }
    })
}



export function queryBookDetail (id) {
    return instance({
        url: '/v1/book/queryBookDetail/' + id,
        method: 'get',
    })
}

export function queryBookIndexAbout (bookId, BookIndexId) {
    return instance({
        url: '/v1/book/queryBookIndexAbout',
        method: 'get',
        params: {
            book_id: bookId,
            book_index_id: BookIndexId
        }
    })
}

// 查看书籍评论列表
export function listCommentByPage (bookId, page = 1, pageSize = 10) {
    return instance({
        url: '/v1/book/listCommentByPage',
        method: 'get',
        params: {
            book_id: bookId,
            page: page,
            page_size: pageSize
        }
    })
}

// 同类书籍推荐
export function listRecBookByCatId (bookId, catId) {
    return instance({
        url: '/v1/book/listRecBookByCatId',
        method: 'get',
        params: {
            book_id: bookId,
            cat_id: catId,
        }
    })
}

// 添加访问量
export function addVisitCount (bookId) {
    return instance({
        url: '/v1/book/addVisitCount/' + bookId,
        method: 'post',
        data: {
            book_id: bookId,
        }
    })
}

// 添加访问量
export function getBookContent (bookId, indexId) {
    return instance({
        url: `/v1/book/content/${bookId}/${indexId}`,
        method: 'get'
    })
}

// 获取书籍访问,页码, 不传pageSize表示查询所有目录列表
export function getIndexList (bookId, page = 1, pageSize = 0) {
    return instance({
        url: `/v1/book/queryIndexList`,
        method: 'get',
        params: {
            book_id: bookId,
            page,
            page_size: pageSize
        }
    })
}

// 获取分类列表
export function listBookCategory () {
    return instance({
        url: `/v1/book/listBookCategory`,
        method: 'get',
    })
}

export function searchBookPage(params) {
    return instance({
        url: `/v1/book/searchByPage`,
        method: 'get',
        params
    })
}
