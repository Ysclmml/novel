// 登录
import instance from "@/utils/request";

// 首页新闻
export function listIndexNews (params) {
    return instance({
        url: '/v1/news/listIndexNews',
        method: 'get',
        params
    })
}
