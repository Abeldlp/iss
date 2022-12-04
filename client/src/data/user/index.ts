import Api from "../Api";
import { User } from "../../types/entities/user";
import { ErrorHandler } from "../ErrorHandling";

export class UserApi {
  private api: Api;
  private errorHandling: ErrorHandler;

  constructor() {
    this.api = new Api();
    this.errorHandling = new ErrorHandler();
  }

  async subscribeUser(payload: any): Promise<User> {
    const data = await this.api.post<User>("/users", payload);

    if (this.errorHandling.isError(data)) {
      throw data;
    }

    return data;
  }
}
