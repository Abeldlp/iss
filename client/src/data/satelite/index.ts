import Api from "../Api";
import { SateliteLocation } from "../../types/entities/sateliteLocation";
import { ErrorHandler } from "../ErrorHandling";

export class SateliteApi {
  private api: Api;
  private errorHandling: ErrorHandler;

  constructor() {
    this.api = new Api();
    this.errorHandling = new ErrorHandler();
  }

  async getAll(params?: any): Promise<SateliteLocation[]> {
    const data = await this.api.get<SateliteLocation[]>("/", params);

    if (this.errorHandling.isError(data)) {
      throw data;
    }

    return data;
  }

  async getAggregated(params?: any): Promise<SateliteLocation[]> {
    const data = await this.api.get<SateliteLocation[]>("/grouped", params);

    if (this.errorHandling.isError(data)) {
      throw data;
    }

    return data;
  }
}
