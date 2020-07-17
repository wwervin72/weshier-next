import Cookies from 'js-cookie';
import {LoginToken} from './variables';

// 禁用掉cookies默认的转义
let customCookie = Cookies.withConverter({
	read: function (value) {
		return value;
	},
	write: function (value) {
		return value;
	}
});

export function getLoginToken (tokenName = LoginToken) {
	return customCookie.get(tokenName);
}

export function setLoginToken (token, tokenName = LoginToken) {
	return customCookie.set(tokenName, token, {
		expires: 1 * 7,
		path: '/'
	});
}

export function removeLoginToken (tokenName = LoginToken) {
	return customCookie.remove(tokenName, {
		path: '/'
	});
}
