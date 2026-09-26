//  this is made for the purpose of creating custom error messages for the API
class ApiError extends Error {
  constructor(message, statusCode) {
    super(message);
    this.statusCode = statusCode;
  }
}
