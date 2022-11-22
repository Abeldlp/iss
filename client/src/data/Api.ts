import axios, { AxiosError, AxiosResponse } from "axios";
import { ErrorHandler, ErrorInterface } from "./ErrorHandling";

export default class Api {
  //header configuration here
  private url: string;

  constructor() {
    this.url = import.meta.env.VITE_API_URL;
  }

  async get(endpoint: string): Promise<AxiosResponse | ErrorInterface> {
    try {
      const res: AxiosResponse = await axios.get(this.url + endpoint);
      return res.data;
    } catch (error) {
      const err = error as AxiosError;
      return new ErrorHandler(
        err.code as string,
        err.message as string
      ).constructError() as ErrorInterface;
    }
  }

  async post(
    endpoint: string,
    paylaod: any
  ): Promise<AxiosResponse | ErrorInterface> {
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

  async put(
    endpoint: string,
    paylaod: any
  ): Promise<AxiosResponse | ErrorInterface> {
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
