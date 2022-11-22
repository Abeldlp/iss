import { AxiosRequestConfig } from "axios";
import Api from "../Api";

export class SateliteApi {
  private api: Api;

  constructor() {
    this.api = new Api();
  }

  getAll(config?: AxiosRequestConfig) {
    this.api.get("/satelite", config);
  }
}
