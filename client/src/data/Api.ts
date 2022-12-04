import axios, { AxiosError, AxiosResponse } from "axios";
import { ErrorHandler, ErrorInterface } from "./ErrorHandling";
import MockDataReader from "../data/mock/mockDataReader";
import AxiosAdapter from "axios-mock-adapter";

export default class Api {
  //Headers configuration here
  private url: string;
  private mock =
    import.meta.env.MODE === "test" ? new AxiosAdapter(axios) : null;

  constructor(url = import.meta.env.VITE_API_URL) {
    this.url = url;
  }

  async get<T>(endpoint: string, params = {}): Promise<T | ErrorInterface> {
    try {
      let adapter;
      if (this.mock) {
        adapter = new AxiosAdapter(axios);
        adapter
          .onGet(this.url + endpoint, { params: params })
          .reply(200, MockDataReader.get(endpoint, params));
      }

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
