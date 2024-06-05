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
    const urlParams = new URLSearchParams(window.location.search)
    return urlParams.get(name)
}