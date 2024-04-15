import { h, render } from 'vue'
import toast from './toast.vue'

let lastVM = {}

function Toast(options) {
    const container = document.createDocumentFragment()
    const vm = h(toast, {
        ...options,
        onDestroy() {
            render(null, container)
        }
    })

    if (Object.keys(lastVM).length !== 0) {
        lastVM.component.props.onDestroy()
        lastVM = {}
    }

    render(vm, container)
    document.body.appendChild(container)
    lastVM = vm

}
export default Toast
