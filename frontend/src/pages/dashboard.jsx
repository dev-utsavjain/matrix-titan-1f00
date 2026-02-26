import { useState, useEffect } from 'react';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function Dashboard() {
  const [user, setUser] = useState(null);
  const [posts, setPosts] = useState([]);
  const [stats, setStats] = useState({
    totalPosts: 0,
    totalViews: 0,
    totalComments: 0,
    accountAge: 0
  });
  const [recentActivity, setRecentActivity] = useState([]);

  useEffect(() => {
    // TODO: connect API endpoint using src/config/api.js
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.PROFILE}`)
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.USER_POSTS}`)
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.USER_STATS}`)
    
    // Mock data
    setUser({
      name: "John Doe",
      email: "john@example.com",
      avatar: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=100&h=100&fit=crop&auto=format"
    });

    setPosts([
      {
        id: 1,
        title: "The Future of Web Development",
        status: "published",
        views: 1250,
        publishDate: "2024-01-15",
        comments: 24
      },
      {
        id: 2,
        title: "Building Scalable Applications",
        status: "draft",
        views: 0,
        publishDate: null,
        comments: 0
      },
      {
        id: 3,
        title: "The Art of Clean Code",
        status: "published",
        views: 890,
        publishDate: "2024-01-10",
        comments: 15
      },
      {
        id: 4,
        title: "Getting Started with React Hooks",
        status: "published",
        views: 2100,
        publishDate: "2024-01-08",
        comments: 32
      }
    ]);

    setStats({
      totalPosts: 12,
      totalViews: 8540,
      totalComments: 156,
      accountAge: 180
    });

    setRecentActivity([
      {
        id: 1,
        type: "view",
        post: "The Future of Web Development",
        user: "Sarah Johnson",
        time: "2 hours ago"
      },
      {
        id: 2,
        type: "comment",
        post: "Building Scalable Applications",
        user: "Mike Chen",
        time: "4 hours ago"
      },
      {
        id: 3,
        type: "view",
        post: "The Art of Clean Code",
        user: "Emily Davis",
        time: "6 hours ago"
      },
      {
        id: 4,
        type: "comment",
        post: "Getting Started with React Hooks",
        user: "Lisa Wang",
        time: "1 day ago"
      }
    ]);
  }, []);

  const handleDeletePost = (postId) => {
    if (window.confirm('Are you sure you want to delete this post?')) {
      setPosts(posts.filter(post => post.id !== postId));
    }
  };

  const handleEditPost = (postId) => {
    // Navigate to edit post page
    console.log('Edit post:', postId);
  };

  const getStatusBadge = (status) => {
    const styles = {
      published: 'bg-green-100 text-green-800',
      draft: 'bg-gray-100 text-gray-800'
    };
    return (
      <span className={`px-3 py-1 rounded-full text-sm font-medium ${styles[status]}`}>
        {status.charAt(0).toUpperCase() + status.slice(1)}
      </span>
    );
  };

  const getActivityIcon = (type) => {
    const icons = {
      view: 'Eye',
      comment: 'MessageCircle'
    };
    const colors = {
      view: 'text-blue-600',
      comment: 'text-green-600'
    };
    return (
      <div className={`w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center ${colors[type]}`}>
        <Icon name={icons[type]} className="w-4 h-4" />
      </div>
    );
  };

  if (!user) return <div className="min-h-screen flex items-center justify-center">Loading...</div>;

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
      <div className="container mx-auto px-4 py-12">
        {/* Dashboard Header */}
        <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6 mb-8">
          <div className="flex items-center gap-4">
            <img
              src={user.avatar}
              alt={user.name}
              className="w-16 h-16 rounded-full object-cover"
              onError={(e) => {
                e.target.onerror = null;
                e.target.src = 'https://placehold.co/100x100/1a1a2e/eaeaea?text=User';
              }}
            />
            <div>
              <h1 className="text-2xl font-bold text-gray-900">Welcome back, {user.name}!</h1>
              <p className="text-gray-600">{user.email}</p>
            </div>
          </div>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-600">Total Posts</p>
                <p className="text-3xl font-bold text-gray-900">{stats.totalPosts}</p>
              </div>
              <div className="w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center">
                <Icon name="FileText" className="w-6 h-6 text-blue-600" />
              </div>
            </div>
          </div>
          
          <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-600">Total Views</p>
                <p className="text-3xl font-bold text-gray-900">{stats.totalViews.toLocaleString()}</p>
              </div>
              <div className="w-12 h-12 bg-green-100 rounded-xl flex items-center justify-center">
                <Icon name="Eye" className="w-6 h-6 text-green-600" />
              </div>
            </div>
          </div>
          
          <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-600">Comments</p>
                <p className="text-3xl font-bold text-gray-900">{stats.totalComments}</p>
              </div>
              <div className="w-12 h-12 bg-purple-100 rounded-xl flex items-center justify-center">
                <Icon name="MessageCircle" className="w-6 h-6 text-purple-600" />
              </div>
            </div>
          </div>
          
          <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-gray-600">Account Age</p>
                <p className="text-3xl font-bold text-gray-900">{stats.accountAge}d</p>
              </div>
              <div className="w-12 h-12 bg-orange-100 rounded-xl flex items-center justify-center">
                <Icon name="Calendar" className="w-6 h-6 text-orange-600" />
              </div>
            </div>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Posts Management */}
          <div className="lg:col-span-2">
            <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
              <div className="flex items-center justify-between mb-6">
                <h2 className="text-xl font-bold text-gray-900">Your Posts</h2>
                <button className="bg-gradient-to-r from-blue-600 to-purple-600 text-white px-4 py-2 rounded-xl font-medium hover:shadow-lg transform hover:scale-[1.02] transition-all duration-300 flex items-center gap-2">
                  <Icon name="Plus" className="w-4 h-4" />
                  New Post
                </button>
              </div>
              
              <div className="overflow-x-auto">
                <table className="w-full">
                  <thead>
                    <tr className="border-b border-gray-200">
                      <th className="text-left py-3 px-4 font-medium text-gray-700">Title</th>
                      <th className="text-left py-3 px-4 font-medium text-gray-700">Status</th>
                      <th className="text-left py-3 px-4 font-medium text-gray-700">Views</th>
                      <th className="text-left py-3 px-4 font-medium text-gray-700">Date</th>
                      <th className="text-left py-3 px-4 font-medium text-gray-700">Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {posts.map((post) => (
                      <tr key={post.id} className="border-b border-gray-100 hover:bg-gray-50/50 transition-colors">
                        <td className="py-3 px-4">
                          <div className="font-medium text-gray-900">{post.title}</div>
                          <div className="text-sm text-gray-500">{post.comments} comments</div>
                        </td>
                        <td className="py-3 px-4">{getStatusBadge(post.status)}</td>
                        <td className="py-3 px-4 text-gray-900">{post.views.toLocaleString()}</td>
                        <td className="py-3 px-4 text-gray-500">
                          {post.publishDate || 'Draft'}
                        </td>
                        <td className="py-3 px-4">
                          <div className="flex gap-2">
                            <button
                              onClick={() => handleEditPost(post.id)}
                              className="text-blue-600 hover:text-blue-700 p-1"
                            >
                              <Icon name="Edit" className="w-4 h-4" />
                            </button>
                            <button
                              onClick={() => handleDeletePost(post.id)}
                              className="text-red-600 hover:text-red-700 p-1"
                            >
                              <Icon name="Trash2" className="w-4 h-4" />
                            </button>
                            <Link
                              to={`/blog/${post.id}`}
                              className="text-gray-600 hover:text-gray-700 p-1"
                            >
                              <Icon name="ExternalLink" className="w-4 h-4" />
                            </Link>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          </div>

          {/* Sidebar */}
          <div className="space-y-8">
            {/* Quick Actions */}
            <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
              <h3 className="text-lg font-bold text-gray-900 mb-4">Quick Actions</h3>
              <div className="space-y-3">
                <button className="w-full bg-gradient-to-r from-blue-600 to-purple-600 text-white py-3 rounded-xl font-medium hover:shadow-lg transform hover:scale-[1.02] transition-all duration-300 flex items-center justify-center gap-2">
                  <Icon name="Edit3" className="w-4 h-4" />
                  Write New Post
                </button>
                <button className="w-full bg-white border border-gray-200 text-gray-700 py-3 rounded-xl font-medium hover:bg-gray-50 transition-colors flex items-center justify-center gap-2">
                  <Icon name="User" className="w-4 h-4" />
                  Edit Profile
                </button>
                <button className="w-full bg-white border border-gray-200 text-gray-700 py-3 rounded-xl font-medium hover:bg-gray-50 transition-colors flex items-center justify-center gap-2">
                  <Icon name="Settings" className="w-4 h-4" />
                  Settings
                </button>
              </div>
            </div>

            {/* Recent Activity */}
            <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6">
              <h3 className="text-lg font-bold text-gray-900 mb-4">Recent Activity</h3>
              <div className="space-y-4">
                {recentActivity.map((activity) => (
                  <div key={activity.id} className="flex items-start gap-3">
                    {getActivityIcon(activity.type)}
                    <div className="flex-1">
                      <p className="text-sm text-gray-900">
                        <span className="font-medium">{activity.user}</span>
                        {activity.type === 'view' ? ' viewed ' : ' commented on '}
                        <span className="font-medium">{activity.post}</span>
                      </p>
                      <p className="text-xs text-gray-500">{activity.time}</p>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}