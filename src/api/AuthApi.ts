export default class AuthApi {
  static user = { username: '', password: '' }
  static token = import.meta.env.VITE_X_TOKEN

  static auth() {
    const user = sessionStorage.getItem('user')
    return user ? JSON.parse(user) : null
  }

  static login(username: string, password: string) {
    if (!this.auth()) {
      this.user = { username, password }
      sessionStorage.setItem('user', JSON.stringify({ username, password }))
    }
  }

  static logout() {
    this.user = { username: '', password: '' }
    sessionStorage.removeItem('user')
  }
}
