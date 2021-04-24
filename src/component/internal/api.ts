import * as Api from '../api'

export class NekoApi {
  api_configuration = new Api.Configuration({
    basePath: location.href.replace(/\/+$/, ''),
    baseOptions: { withCredentials: true },
  })

  public setUrl(url: string) {
    this.api_configuration.basePath = url.replace(/\/+$/, '')
  }

  public setToken(token: string) {
    this.api_configuration.accessToken = token
  }

  get url(): string {
    return this.api_configuration.basePath || location.href.replace(/\/+$/, '')
  }

  get session(): SessionApi {
    return new Api.SessionApi(this.api_configuration)
  }

  get room(): RoomApi {
    return new Api.RoomApi(this.api_configuration)
  }

  get members(): MembersApi {
    return new Api.MembersApi(this.api_configuration)
  }
}

export type SessionApi = Api.SessionApi
export type RoomApi = Api.RoomApi
export type MembersApi = Api.MembersApi
