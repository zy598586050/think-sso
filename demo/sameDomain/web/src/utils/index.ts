// 获取Cookie
export const getCookie = (name: string) => {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
}

// 删除Cookie
export const deleteCookie = (name: string) => {
    // 设置该cookie的过期时间为过去的一个时间点
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

// 从URL中获取参数
export const getQueryParam = (name: string) => {
    // 定义正则表达式用于匹配参数
    const regex = new RegExp(`[?&]${name}=([^&#]*)`, 'i')
    
    // 优先从 URL 的查询字符串中查找参数
    const searchResults = window.location.search.match(regex)
    if (searchResults) {
        return decodeURIComponent(searchResults[1])
    }
    
    // 如果查询字符串中没有找到参数，再从哈希片段中查找
    const hashResults = window.location.hash.match(regex)
    if (hashResults) {
        return decodeURIComponent(hashResults[1])
    }
    
    // 如果两个位置都没有找到参数，返回 null
    return null
}