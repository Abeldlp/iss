export interface ErrorInterface {
  code: string;
  message: string;
}

export class ErrorHandler {
  private errorCode: string;
  private errorMessage: string;

  constructor(errorCode: string, errorMessage: string) {
    this.errorCode = errorCode;
    this.errorMessage = errorMessage;
  }

  constructError(): ErrorInterface {
    return {
      code: this.errorCode,
      message: this.errorMessage,
    };
  }

  throwError(): Error {
    throw new Error(this.errorMessage);
  }
}
