import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function BlogPosts() {
  const [posts, setPosts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [filteredPosts, setFilteredPosts] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [sortBy, setSortBy] = useState('newest');
  const [currentPage, setCurrentPage] = useState(1);
  const postsPerPage = 12;

  useEffect(() => {
    // TODO: connect API endpoint using src/config/api.js
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.POSTS}`)
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.CATEGORIES}`)
    
    // Mock data
    const mockPosts = [
      {
        id: 1,
        title: "The Future of Web Development",
        excerpt: "Exploring the latest trends and technologies shaping the future of web development...",
        author: "Sarah Johnson",
        publishDate: "2024-01-15",
        readTime: "5 min read",
        category: "Technology",
        views: 1250,
        image: "https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 2,
        title: "Building Scalable Applications",
        excerpt: "Best practices for creating applications that can grow with your user base...",
        author: "Mike Chen",
        publishDate: "2024-01-12",
        readTime: "8 min read",
        category: "Development",
        views: 980,
        image: "https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 3,
        title: "The Art of Clean Code",
        excerpt: "Writing code that is not only functional but also maintainable and elegant...",
        author: "Emily Davis",
        publishDate: "2024-01-10",
        readTime: "6 min read",
        category: "Best Practices",
        views: 1450,
        image: "https://images.unsplash.com/photo-1518770660439-4636190af475?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 4,
        title: "Getting Started with React Hooks",
        excerpt: "A comprehensive guide to understanding and using React Hooks effectively...",
        author: "John Smith",
        publishDate: "2024-01-08",
        readTime: "7 min read",
        category: "React",
        views: 2100,
        image: "https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 5,
        title: "CSS Grid vs Flexbox",
        excerpt: "When to use CSS Grid and when to use Flexbox for your layouts...",
        author: "Lisa Wang",
        publishDate: "2024-01-06",
        readTime: "4 min read",
        category: "CSS",
        views: 890,
        image: "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 6,
        title: "JavaScript Performance Tips",
        excerpt: "Practical tips to improve the performance of your JavaScript applications...",
        author: "David Brown",
        publishDate: "2024-01-04",
        readTime: "9 min read",
        category: "JavaScript",
        views: 1680,
        image: "https://images.unsplash.com/photo-1460925895917-afdab827c52f?w=400&h=300&fit=crop&auto=format"
      }
    ];
    
    setPosts(mockPosts);
    setFilteredPosts(mockPosts);
    setCategories(["All", "Technology", "Development", "Best Practices", "React", "CSS", "JavaScript"]);
  }, []);

  useEffect(() => {
    let filtered = posts;
    
    if (selectedCategory !== 'all') {
      filtered = filtered.filter(post => post.category === selectedCategory);
    }
    
    if (searchTerm) {
      filtered = filtered.filter(post => 
        post.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
        post.excerpt.toLowerCase().includes(searchTerm.toLowerCase()) ||
        post.author.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }
    
    if (sortBy === 'oldest') {
      filtered = [...filtered].sort((a, b) => new Date(a.publishDate) - new Date(b.publishDate));
    } else if (sortBy === 'popular') {
      filtered = [...filtered].sort((a, b) => b.views - a.views);
    } else {
      filtered = [...filtered].sort((a, b) => new Date(b.publishDate) - new Date(a.publishDate));
    }
    
    setFilteredPosts(filtered);
    setCurrentPage(1);
  }, [posts, searchTerm, selectedCategory, sortBy]);

  const totalPages = Math.ceil(filteredPosts.length / postsPerPage);
  const startIndex = (currentPage - 1) * postsPerPage;
  const currentPosts = filteredPosts.slice(startIndex, startIndex + postsPerPage);

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
      <div className="container mx-auto px-4 py-12">
        {/* Page Header */}
        <div className="text-center mb-12">
          <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-4">All Posts</h1>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto">
            Discover amazing stories from our community of writers and creators
          </p>
        </div>

        {/* Filter Bar */}
        <div className="bg-white/70 backdrop-blur-lg rounded-2xl shadow-xl border border-white/20 p-6 mb-8">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="relative">
              <Icon name="Search" className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
              <input
                type="text"
                placeholder="Search posts..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="w-full pl-10 pr-4 py-3 bg-white/50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-300"
              />
            </div>
            
            <select
              value={selectedCategory}
              onChange={(e) => setSelectedCategory(e.target.value)}
              className="w-full px-4 py-3 bg-white/50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-300"
            >
              {categories.map(category => (
                <option key={category} value={category.toLowerCase()}>{category}</option>
              ))}
            </select>
            
            <select
              value={sortBy}
              onChange={(e) => setSortBy(e.target.value)}
              className="w-full px-4 py-3 bg-white/50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-300"
            >
              <option value="newest">Newest First</option>
              <option value="oldest">Oldest First</option>
              <option value="popular">Most Popular</option>
            </select>
          </div>
        </div>

        {/* Posts Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-12">
          {currentPosts.map((post) => (
            <article key={post.id} className="bg-white rounded-2xl shadow-lg overflow-hidden hover:shadow-xl transition-all duration-300 hover:scale-[1.02]">
              <img
                src={post.image}
                alt={post.title}
                className="w-full h-48 object-cover"
                onError={(e) => {
                  e.target.onerror = null;
                  e.target.src = 'https://placehold.co/400x300/1a1a2e/eaeaea?text=Image';
                }}
              />
              <div className="p-6">
                <div className="flex items-center gap-2 mb-3">
                  <span className="px-3 py-1 bg-blue-100 text-blue-800 text-sm rounded-full font-medium">
                    {post.category}
                  </span>
                  <span className="text-gray-500 text-sm">{post.views} views</span>
                </div>
                <h2 className="text-xl font-bold text-gray-900 mb-3">{post.title}</h2>
                <p className="text-gray-600 mb-4 line-clamp-2">{post.excerpt}</p>
                <div className="flex items-center justify-between text-sm text-gray-500 mb-4">
                  <span>{post.author}</span>
                  <span>{post.publishDate}</span>
                </div>
                <div className="flex items-center justify-between">
                  <span className="text-sm text-gray-500">{post.readTime}</span>
                  <Link
                    to={`/blog/${post.id}`}
                    className="inline-flex items-center gap-1 text-blue-600 hover:text-blue-700 font-medium"
                  >
                    Read More
                    <Icon name="ArrowRight" className="w-4 h-4" />
                  </Link>
                </div>
              </div>
            </article>
          ))}
        </div>

        {/* No Results */}
        {filteredPosts.length === 0 && (
          <div className="text-center py-12">
            <Icon name="SearchX" className="w-16 h-16 text-gray-400 mx-auto mb-4" />
            <p className="text-xl text-gray-600">No posts found matching your criteria.</p>
          </div>
        )}

        {/* Pagination */}
        {totalPages > 1 && (
          <div className="flex justify-center gap-2">
            <button
              onClick={() => setCurrentPage(prev => Math.max(prev - 1, 1))}
              disabled={currentPage === 1}
              className="px-4 py-2 bg-white border border-gray-200 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 transition-colors"
            >
              <Icon name="ChevronLeft" className="w-5 h-5" />
            </button>
            
            {Array.from({ length: totalPages }, (_, i) => i + 1).map(page => (
              <button
                key={page}
                onClick={() => setCurrentPage(page)}
                className={`px-4 py-2 rounded-lg transition-colors ${
                  currentPage === page
                    ? 'bg-blue-600 text-white'
                    : 'bg-white border border-gray-200 hover:bg-gray-50'
                }`}
              >
                {page}
              </button>
            ))}
            
            <button
              onClick={() => setCurrentPage(prev => Math.min(prev + 1, totalPages))}
              disabled={currentPage === totalPages}
              className="px-4 py-2 bg-white border border-gray-200 rounded-lg disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-50 transition-colors"
            >
              <Icon name="ChevronRight" className="w-5 h-5" />
            </button>
          </div>
        )}
      </div>
    </div>
  );
}