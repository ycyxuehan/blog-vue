/* jshint esversion: 6 */
import axios from 'axios'

let base = 'api/v1'

export const GetCategories = (params) => { return axios.get(`${base}/category`, {params: params}).then(res => res.data) }
export const GetCategory = name => { return axios.get(`${base}/category/${name}`).then(res => res.data) }
export const AddCategory = (params) => { return axios.post(`${base}/category`, params).then(res => res.data) }
export const SetCategory = (name, params) => { return axios.put(`${base}/category/${name}`, params).then(res => res.data) }
export const DelCategory = name => { return axios.delete(`${base}/category/${name}`).then(res => res.data) }

export const GetArticles = (name, params) => { return axios.get(`${base}/article/${name}`, {params: params}).then(res => res.data) }
export const GetArticle = name => { return axios.get(`${base}/article/${name}`).then(res => res.data) }
export const AddArticle = (name, params) => { return axios.post(`${base}/article/${name}`, params).then(res => res.data) }
export const SetArticle = (name, params) => { return axios.put(`${base}/article/${name}`, params).then(res => res.data) }
export const DelArticle = name => { return axios.delete(`${base}/article/${name}`).then(res => res.data) }

export const Login = params => { return axios.post(`${base}/user/login`, params).then(res => res.data) }
export const VerifySession = params => { return axios.get(`${base}/user/session`, {params: params}).then(res => res.data) }
