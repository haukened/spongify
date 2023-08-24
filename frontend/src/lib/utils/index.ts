export const isValidURL = (url: string):boolean => {
    const r = new RegExp('^\/[a-zA-Z0-9\/\\.\?=&_#%]*$')
    return r.test(url)
}