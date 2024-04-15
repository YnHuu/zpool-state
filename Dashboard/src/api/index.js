import { request } from '@/utils/request'

export function loginApi(data) {
    return request({
        url: '/api/login',
        method: 'post',
        data
    })
}

export function getZFS() {
    return request({
        url: '/api/zfs',
        method: 'get',
    })
}

export function getDisk() {
    return request({
        url: '/api/disk',
        method: 'get',
    })
}

export function getCron(uid) {
    return request({
        url: '/api/crontab/' + uid,
        method: 'get',
    })
}

export function setCron(uid, data) {
    return request({
        url: '/api/crontab/' + uid,
        method: 'post',
        data
    })
}

export function getWechat() {
    return request({
        url: '/api/wechat',
        method: 'get',
    })
}

export function setWechat(data) {
    return request({
        url: '/api/wechat',
        method: 'post',
        data
    })
}

export function testWechat() {
    return request({
        url: '/api/wechat/test',
        method: 'get',
    })
}

export function rebootApi() {
    return request({
        url: '/api/reboot',
        method: 'get',
    })
}