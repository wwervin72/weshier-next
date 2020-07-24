import React from 'react'
import {Link} from 'react-router-dom'
import PropTypes from 'prop-types'
import './index.scss'

export default function Avatar ({user, width = 30, height = 30, src, alt = 'weshier'}) {
	return (
		user ?
		<Link className="avatar" style={{
			width: `${width}px`,
			height: `${height}px`
		}} to={`/u/${user.id}`}>
			<img width="100%" height="100%" alt={alt || user.nickName}
				src={src || user.avatar || require('../../assets/img/avatar.png')}></img>
		</Link> :
		<img className="avatar" width={width} height={height} alt={alt}
			src={src || require('../../assets/img/avatar.png')}></img>
	)
}

Avatar.propTypes = {
	user: PropTypes.object,
	width: PropTypes.number,
	height: PropTypes.number,
	src: PropTypes.string,
	alt: PropTypes.string,
}

