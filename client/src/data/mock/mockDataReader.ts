import { MockData } from "./mockData";

export default class MockDataReader {
  private static consoleStyle = "color: #287cb9";

  static get(endpoint: string, config: any): Record<never, never> | undefined {
    console.groupCollapsed(`%cread mock for ${endpoint}`, this.consoleStyle);
    console.groupEnd();
    const subresource = endpoint.split("/")[1];

    return subresource
      ? MockData[subresource].get(config)
      : MockData.get(config);
  }
}
