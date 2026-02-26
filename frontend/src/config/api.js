const API_BASE_URL = "";

export const API_ENDPOINTS = {
  GET_POSTS: "/api/posts",
  GET_FEATURED_POSTS: "/api/posts/featured",
  GET_RECENT_POSTS: "/api/posts/recent",
  GET_POST_BY_SLUG: "/api/posts/:slug",
  GET_RELATED_POSTS: "/api/posts/related/:id",
  GET_CATEGORIES: "/api/categories",
  POST_SIGNUP: "/api/auth/signup",
  POST_LOGIN: "/api/auth/login",
  GET_CHECK_USERNAME: "/api/auth/check-username",
  GET_USER_PROFILE: "/api/users/profile",
  GET_USER_POSTS: "/api/users/posts",
  GET_USER_STATS: "/api/users/stats"
};

export default { API_BASE_URL, API_ENDPOINTS };