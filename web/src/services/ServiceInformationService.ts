import http from '../http-commons'

export class ServiceInformationService {
  get () {
    return http.get('/v1/self/info')
  }
}
