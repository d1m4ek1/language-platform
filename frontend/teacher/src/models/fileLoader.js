class TeacherFileLoader {
  static async fileLoad(body, type, ...params) {
    const paramItems = [`upload_type=${type}`, ...params]

    return await fetch(`/teacher/upload-file?${paramItems.join("&")}`, {
      method: "POST",
      body
    })
  }
}

class APIFileLoader extends TeacherFileLoader {}

export default APIFileLoader