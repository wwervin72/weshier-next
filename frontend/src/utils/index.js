/**
 * 根据请求体中返回的 code 来判断，提示消息的状态
 * @param {*} code
 */
export function parseMessageStatusByResCode(code) {
	let status = 'success';
	if (code !== 200) {
		status = 'error';
	}
	return status;
}
