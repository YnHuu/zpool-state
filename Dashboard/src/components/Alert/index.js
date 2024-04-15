import { h, render } from 'vue'
import alert from './alert.vue'

const instances = []

function Alert(options) {
    let top = 30;

    instances.forEach(vm => {
        top += vm.el.offsetHeight + 48
    });

    const container = document.createDocumentFragment()

    const vm = h(alert, {
        ...options,
        top,
        onClose() {
            close(vm)
        },
        onDestroy() {
            render(null, container)
        }
    })

    render(vm, container)
    document.body.appendChild(container)
    instances.push(vm)
}

function close(vm) {
    const index = instances.findIndex(ins => ins === vm)
    if (index === -1) {
        return;
    }

    instances.splice(index, 1)

    for (let start = index; start < instances.length; start++) {
        const cpn = instances[start].component
        cpn.props.top -= vm.el.offsetHeight + 48
    }
}

export default Alert
