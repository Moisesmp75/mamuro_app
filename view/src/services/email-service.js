const main_url = "http://localhost:3000";
const url = `${main_url}/api/v1/mail`

export class MailService {
  constructor() {
    this.requestOption = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    }
  }
  async search_data(request) {
    try {
      this.requestOption["body"] = JSON.stringify(request)
      const response = await fetch(`${url}/search`, this.requestOption)
      if(!response.ok)
        throw new Error(response.status)
      const { success, message, data, meta } = await response.json()
      return {
        success,
        message,
        data,
        meta
      }
    } catch (error) {
      console.log(error)
      throw error
    }
  }
}