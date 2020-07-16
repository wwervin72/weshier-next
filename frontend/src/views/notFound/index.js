import React from 'react'
import {Link} from 'react-router-dom'

export default function NotFound() {
	return (
		<p>未找到页面<Link to="/">返回首页</Link></p>
	)
}
