import { createApp } from 'vue'
import loading from './Loading.vue'

let app = {}
const Loading = {
    show() {
        const container = document.createDocumentFragment()
        app = createApp(loading)
        app.mount(container)
        document.body.appendChild(container)
    },
    hide() {
        app.unmount()
    }
}
export default Loading