import { createApp } from 'vue'
import Dialog from './dialog.vue'

function MessageBox(options) {
    return new Promise((resolve) => {
        const app = createApp(Dialog, {
            ...options,
            onDestroy: () => {
                app.unmount()
            },
            onOK: () => {
                resolve()
                app.unmount()
            },
        })

        const container = document.createDocumentFragment()
        app.mount(container)
        document.body.appendChild(container)

    })
}


export default MessageBox