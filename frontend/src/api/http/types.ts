import type { InternalAxiosRequestConfig, AxiosRequestConfig } from 'axios'

export interface InternalRequestConfig extends InternalAxiosRequestConfig {
  needToken?: boolean
}

export interface RequestConfig extends AxiosRequestConfig {
  needToken?: boolean
}
