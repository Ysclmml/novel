// 全局过滤器

import {dateFormat} from "@/utils"
import Vue from 'vue'

// 根据字符串'2021:01:02 10:01:01'来格式化时间
function timeFormat(time, format) {
    const d = new Date(time)
    return dateFormat(d.getTime(), format)
}

export function registerFilters() {
    Vue.filter('timeFormat', timeFormat)
}
