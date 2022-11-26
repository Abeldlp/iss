import { describe, it, expect } from "vitest";
import { ErrorHandler } from "../data/ErrorHandling";

describe("Error handling test case", () => {
  it("can instantiate error handler and return code", () => {
    const errorHandler = new ErrorHandler("200", "Status OK");

    const errorCode = errorHandler.constructError();

    expect(errorCode.code).toBe("200");
  });

  it("can instantiate error handler and return message", () => {
    const errorHandler = new ErrorHandler("200", "Status OK");

    const errorCode = errorHandler.constructError();

    expect(errorCode.message).toBe("Status OK");
  });

  it("can instantiate error handler and recognize error response", () => {
    const errorHandler = new ErrorHandler("200", "Status OK");

    const errorCode = errorHandler.constructError();

    expect(errorHandler.isError(errorCode)).toBe(true);
  });
});
