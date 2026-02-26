import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function Home() {
  const [featuredPosts, setFeaturedPosts] = useState([]);
  const [recentPosts, setRecentPosts] = useState([]);
  const [currentSlide, setCurrentSlide] = useState(0);

  useEffect(() => {
    // TODO: connect API endpoint using src/config/api.js
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.FEATURED_POSTS}`)
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.RECENT_POSTS}`)
    
    // Mock data
    setFeaturedPosts([
      {
        id: 1,
        title: "The Future of Web Development",
        excerpt: "Exploring the latest trends and technologies shaping the future of web development...",
        author: "Sarah Johnson",
        publishDate: "2024-01-15",
        readTime: "5 min read",
        image: "https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=800&h=400&fit=crop&auto=format"
      },
      {
        id: 2,
        title: "Building Scalable Applications",
        excerpt: "Best practices for creating applications that can grow with your user base...",
        author: "Mike Chen",
        publishDate: "2024-01-12",
        readTime: "8 min read",
        image: "https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=800&h=400&fit=crop&auto=format"
      },
      {
        id: 3,
        title: "The Art of Clean Code",
        excerpt: "Writing code that is not only functional but also maintainable and elegant...",
        author: "Emily Davis",
        publishDate: "2024-01-10",
        readTime: "6 min read",
        image: "https://images.unsplash.com/photo-1518770660439-4636190af475?w=800&h=400&fit=crop&auto=format"
      }
    ]);

    setRecentPosts([
      {
        id: 4,
        title: "Getting Started with React Hooks",
        excerpt: "A comprehensive guide to understanding and using React Hooks effectively...",
        author: "John Smith",
        publishDate: "2024-01-08",
        readTime: "7 min read",
        image: "https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 5,
        title: "CSS Grid vs Flexbox",
        excerpt: "When to use CSS Grid and when to use Flexbox for your layouts...",
        author: "Lisa Wang",
        publishDate: "2024-01-06",
        readTime: "4 min read",
        image: "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 6,
        title: "JavaScript Performance Tips",
        excerpt: "Practical tips to improve the performance of your JavaScript applications...",
        author: "David Brown",
        publishDate: "2024-01-04",
        readTime: "9 min read",
        image: "https://images.unsplash.com/photo-1460925895917-afdab827c52f?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 7,
        title: "Understanding TypeScript",
        excerpt: "Why TypeScript is becoming the standard for large-scale JavaScript projects...",
        author: "Anna Lee",
        publishDate: "2024-01-02",
        readTime: "6 min read",
        image: "https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 8,
        title: "API Design Best Practices",
        excerpt: "How to design RESTful APIs that are intuitive and maintainable...",
        author: "Tom Wilson",
        publishDate: "2023-12-30",
        readTime: "8 min read",
        image: "https://images.unsplash.com/photo-1561998338-13ad7883b20f?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 9,
        title: "Mobile-First Development",
        excerpt: "Why starting with mobile design leads to better user experiences...",
        author: "Maria Garcia",
        publishDate: "2023-12-28",
        readTime: "5 min read",
        image: "https://images.unsplash.com/photo-1452587925148-ce544e77e70d?w=400&h=300&fit=crop&auto=format"
      }
    ]);
  }, []);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % featuredPosts.length);
    }, 5000);
    return () => clearInterval(timer);
  }, [featuredPosts.length]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50">
      {/* Hero Section */}
      <section className="relative py-20 md:py-32 overflow-hidden">
        <div className="absolute inset-0 z-0">
          <img
            src="https://images.unsplash.com/photo-1497366216548-37526070297c?w=1920&h=1080&fit=crop&auto=format"
            alt="Hero background"
            className="w-full h-full object-cover opacity-10"
            onError={(e) => {
              e.target.onerror = null;
              e.target.src = 'https://placehold.co/1920x1080/1a1a2e/eaeaea?text=Image';
            }}
          />
        </div>
        <div className="container mx-auto px-4 relative z-10">
          <div className="text-center max-w-4xl mx-auto">
            <h1 className="text-5xl md:text-7xl lg:text-8xl font-bold text-gray-900 mb-6 animate-fade-in-up">
              Welcome to BlogSphere
            </h1>
            <p className="text-xl md:text-2xl text-gray-600 mb-8 animate-fade-in-up animation-delay-200">
              Share your stories with the world. A modern platform for creators to write, publish, and connect.
            </p>
            <Link
              to="/signup"
              className="inline-flex items-center gap-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-8 py-4 rounded-xl font-medium hover:shadow-lg transform hover:scale-[1.02] transition-all duration-300 animate-fade-in-up animation-delay-400"
            >
              <Icon name="Edit3" className="w-5 h-5" />
              Start Writing Today
            </Link>
          </div>
        </div>
      </section>

      {/* Featured Posts Carousel */}
      <section className="py-16 md:py-24">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl md:text-5xl font-bold text-center text-gray-900 mb-12">Featured Stories</h2>
          <div className="relative max-w-4xl mx-auto">
            <div className="overflow-hidden rounded-2xl shadow-2xl">
              <div className="flex transition-transform duration-500 ease-in-out" style={{ transform: `translateX(-${currentSlide * 100}%)` }}>
                {featuredPosts.map((post) => (
                  <div key={post.id} className="w-full flex-shrink-0">
                    <div className="bg-white rounded-2xl overflow-hidden">
                      <img
                        src={post.image}
                        alt={post.title}
                        className="w-full h-64 md:h-80 object-cover"
                        onError={(e) => {
                          e.target.onerror = null;
                          e.target.src = 'https://placehold.co/800x400/1a1a2e/eaeaea?text=Image';
                        }}
                      />
                      <div className="p-8">
                        <h3 className="text-2xl md:text-3xl font-bold text-gray-900 mb-4">{post.title}</h3>
                        <p className="text-gray-600 mb-4 text-lg">{post.excerpt}</p>
                        <div className="flex items-center justify-between">
                          <div className="flex items-center gap-4 text-sm text-gray-500">
                            <span>{post.author}</span>
                            <span>•</span>
                            <span>{post.publishDate}</span>
                            <span>•</span>
                            <span>{post.readTime}</span>
                          </div>
                          <Link
                            to={`/blog/${post.id}`}
                            className="text-blue-600 hover:text-blue-700 font-medium flex items-center gap-1"
                          >
                            Read More
                            <Icon name="ArrowRight" className="w-4 h-4" />
                          </Link>
                        </div>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
            <div className="flex justify-center gap-2 mt-6">
              {featuredPosts.map((_, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentSlide(index)}
                  className={`w-3 h-3 rounded-full transition-all duration-300 ${
                    currentSlide === index ? 'bg-blue-600' : 'bg-gray-300'
                  }`}
                />
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Recent Posts Grid */}
      <section className="py-16 md:py-24 bg-white/50">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl md:text-5xl font-bold text-center text-gray-900 mb-12">Recent Posts</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {recentPosts.map((post) => (
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
                  <h3 className="text-xl font-bold text-gray-900 mb-3">{post.title}</h3>
                  <p className="text-gray-600 mb-4 line-clamp-2">{post.excerpt}</p>
                  <div className="flex items-center justify-between text-sm text-gray-500 mb-4">
                    <span>{post.author}</span>
                    <span>{post.readTime}</span>
                  </div>
                  <Link
                    to={`/blog/${post.id}`}
                    className="inline-flex items-center gap-2 text-blue-600 hover:text-blue-700 font-medium"
                  >
                    Read More
                    <Icon name="ArrowRight" className="w-4 h-4" />
                  </Link>
                </div>
              </article>
            ))}
          </div>
          <div className="text-center mt-12">
            <Link
              to="/blog"
              className="inline-flex items-center gap-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-8 py-3 rounded-xl font-medium hover:shadow-lg transform hover:scale-[1.02] transition-all duration-300"
            >
              View All Posts
              <Icon name="ArrowRight" className="w-5 h-5" />
            </Link>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 md:py-24">
        <div className="container mx-auto px-4">
          <div className="bg-gradient-to-r from-blue-600 to-purple-600 rounded-3xl p-8 md:p-12 text-center text-white relative overflow-hidden">
            <div className="absolute inset-0 bg-black/10" />
            <div className="relative z-10">
              <h2 className="text-3xl md:text-5xl font-bold mb-4">Ready to Share Your Story?</h2>
              <p className="text-xl mb-8 opacity-90">Join thousands of creators already publishing on BlogSphere</p>
              <Link
                to="/signup"
                className="inline-flex items-center gap-2 bg-white text-blue-600 px-8 py-4 rounded-xl font-bold hover:shadow-xl transform hover:scale-[1.02] transition-all duration-300"
              >
                <Icon name="Rocket" className="w-5 h-5" />
                Start Writing Now
              </Link>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}