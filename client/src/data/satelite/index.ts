import Api from "../Api";

export class SateliteApi {
  private api: Api;

  constructor() {
    this.api = new Api();
  }

  getAll() {
    return this.api.get("/");
  }

  getAggregated() {
    return this.api.get("/grouped");
  }
}
