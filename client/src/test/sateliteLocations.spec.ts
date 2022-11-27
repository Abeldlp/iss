import { describe, it, expect } from "vitest";
import { SateliteApi } from "../data/satelite";

describe("Satelite Locations test case", () => {
  it("get data from api root endpoint", async () => {
    const sateliteApi = new SateliteApi();

    const data = await sateliteApi.getAll();

    expect(data[0].timezone).toBe("Etc/GMT-10");
  });

  it("get data from api grouped endpoint", async () => {
    const sateliteApi = new SateliteApi();

    const data = await sateliteApi.getAggregated();

    expect(data[0].minutes).toBe(3);
  });
});
