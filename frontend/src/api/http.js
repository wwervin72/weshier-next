import axios from "axios";
import qs from "querystring";
import { getLoginToken } from "../utils/cookie";
import { parseMessageStatusByResCode } from "../utils";
import {message} from 'antd'

const CancelToken = axios.CancelToken;
const source = CancelToken.source();
const timeOutMs = 10000;

const instance = axios.create({
	validateStatus: function (status) {
		return status >= 200 && status <= 510;
	},
	withCredentials: true
});

// 错误信息转换
const requestErrorMsg = {
	"Network Error": "网络错误",
	[`timeout of ${timeOutMs}ms exceeded`]: "请求超时"
};

instance.interceptors.response.use(
	(res) => {
		// 当有请求成功 重置isAuthFailed
		let resBody = res.data;
		let msg = resBody.message;
		msg && message[parseMessageStatusByResCode(resBody.code)](msg);
		// 使用 code 来判断该 http 请求是否如预期执行成功
		if (resBody.code === 200) {
			return Promise.resolve(resBody.data);
		} else {
			return Promise.reject(resBody);
		}
	},
	(error) => {
		let msg = requestErrorMsg[error.message];
		message.error(msg);
		return Promise.reject(new Error(msg));
	}
);

instance.interceptors.request.use(
	(config) => {
		if (config.url === "login") {
		} else {
			let token = getLoginToken();
			if (!token) {
				config.data.CancelToken = source.token;
			} else {
				config.headers.token = token;
			}
		}
		if (
			config.method === "post" &&
			config.headers["Content-Type"] ===
				"application/x-www-form-urlencoded"
		) {
			config.data = qs.stringify(config.data);
		}
		return config;
	},
	(err) => {
		return Promise.reject(err);
	}
);

export default ({
	method = "GET",
	baseURL = "",
	url = "",
	data = {},
	params = {},
	headers = {
		"Content-Type": "application/json"
	},
	timeout = timeOutMs
}) => {
	return instance({
		method,
		baseURL,
		url,
		data,
		params,
		headers,
		timeout
	});
};
