import http from '../http-commons'

class ServiceInformationService {
  get () {
    return http.get('/v1/self/info')
  }
}

export default new ServiceInformationService()
