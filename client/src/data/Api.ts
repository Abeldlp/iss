import axios, { AxiosError, AxiosResponse } from "axios";
import { ErrorHandler, ErrorInterface } from "./ErrorHandling";

export default class Api {
  //Headers configuration here
  private url: string;

  constructor() {
    this.url = import.meta.env.VITE_API_URL;
  }

  async get<T>(endpoint: string, params = {}): Promise<T | ErrorInterface> {
    try {
      const res: AxiosResponse = await axios.get(this.url + endpoint, {
        params: params,
      });
      return res.data;
    } catch (error) {
      const err = error as AxiosError;
      return new ErrorHandler(
        err.code as string,
        err.message as string
      ).constructError() as ErrorInterface;
    }
  }

  async post<T>(endpoint: string, paylaod: any): Promise<T | ErrorInterface> {
    try {
      const res: AxiosResponse = await axios.post(this.url + endpoint, paylaod);
      return res.data;
    } catch (error) {
      const err = error as AxiosError;
      return new ErrorHandler(
        err.code as string,
        err.message as string
      ).constructError() as ErrorInterface;
    }
  }

  async put<T>(endpoint: string, paylaod: any): Promise<T | ErrorInterface> {
    try {
      const res: AxiosResponse = await axios.put(this.url + endpoint, paylaod);
      return res.data;
    } catch (error) {
      const err = error as AxiosError;
      return new ErrorHandler(
        err.code as string,
        err.message as string
      ).constructError() as ErrorInterface;
    }
  }

  async delete(endpoint: string): Promise<AxiosResponse | ErrorInterface> {
    try {
      const res: AxiosResponse = await axios.delete(this.url + endpoint);
      return res.data;
    } catch (error) {
      const err = error as AxiosError;
      return new ErrorHandler(
        err.code as string,
        err.message as string
      ).constructError() as ErrorInterface;
    }
  }
}
