const loadFile = (idFile) => {
    return `http://localhost:8080/uploads/${idFile}`
}
 
export default {
    install: (app, options) => {
        app.config.globalProperties.$loadFile = loadFile
        app.provide('$loadFile', loadFile)
    }
}