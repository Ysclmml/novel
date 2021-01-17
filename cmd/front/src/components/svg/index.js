import Vue from 'vue'
import SvgIcon from './SvgIcon'

//eslint-disable-next-line vue/match-component-file-name
Vue.component(`SvgIcon`, SvgIcon)
// requires and returns all modules that match
const requireAll = requireContext => requireContext.keys().map(requireContext)

// import all svg
const req = require.context('@/assets/svg', true, /\.svg$/)
requireAll(req)
