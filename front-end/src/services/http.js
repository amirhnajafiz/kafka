import axios from 'axios'

class HttpService {
    constructor(instance) {
        this.instance = instance
    }

    get(url, config) {
        return this.instance.get(url, config)
    }

    post(url, body, config) {
        return this.instance.post(url, body, config)
    }

    put(url, body, config) {
        return this.instance.put(url, body, config)
    }

    patch(url, body, config) {
        return this.instance.patch(url, body, config)
    }

    delete(url, config) {
        return this.instance.delete(url, config)
    }
}

const httpService = new HttpService(axios.create())
Object.freeze(httpService)
Object.seal(httpService)

export default httpService