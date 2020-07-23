import React from 'react'
import {Link} from 'react-router-dom'
import {connect} from 'react-redux'
// import PropTypes from 'prop-types'

function Avatar ({userInfo}) {
	return (
		<Link to={`/u/${userInfo.id}`}>
			<img alt={userInfo.nickName} src={userInfo.avatar || require('../assets/img/avatar.png')}></img>
		</Link>
	)
}

Avatar.propTypes = {

}

export default connect((state) => ({
	userInfo: state.userInfo
}))(Avatar)
