import Api from "../Api";
import { SateliteLocation } from "../../types/entities/sateliteLocation";
import { ErrorHandler, ErrorInterface } from "../ErrorHandling";

export class SateliteApi {
  private api: Api;
  private errorHandling: ErrorHandler;

  constructor() {
    this.api = new Api();
    this.errorHandling = new ErrorHandler();
  }

  async getAll(): Promise<SateliteLocation[]> {
    const data = await this.api.get<SateliteLocation[]>("/");

    if (this.errorHandling.isError(data)) {
      throw data;
    }

    return data;
  }

  async getAggregated(): Promise<SateliteLocation[]> {
    const data = await this.api.get<SateliteLocation[]>("/grouped");

    if (this.errorHandling.isError(data)) {
      throw data;
    }

    return data;
  }
}
