import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import * as Icons from 'lucide-react';
import Header from '../components/header';
import Footer from '../components/footer';
import { API_BASE_URL, API_ENDPOINTS } from '../config/api';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function BlogPost() {
  const { slug } = useParams();
  const [post, setPost] = useState(null);
  const [relatedPosts, setRelatedPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [commentText, setCommentText] = useState('');
  const [comments, setComments] = useState([
    {
      id: 1,
      author: 'Alex Chen',
      content: 'Great insights on the future of web development! The AI integration section was particularly interesting.',
      timestamp: '2 days ago',
      avatar: '1522071820081-009f0129c71c'
    },
    {
      id: 2,
      author: 'Maria Garcia',
      content: 'This article really helped me understand the direction our industry is heading. Thank you for sharing!',
      timestamp: '1 week ago',
      avatar: '1495474472287-4d71bcdd2085'
    }
  ]);

  useEffect(() => {
    const mockPost = {
      id: 1,
      title: "The Future of Web Development",
      slug: slug,
      content: `<h2>Introduction</h2>
<p>The landscape of web development is evolving at an unprecedented pace. With new technologies emerging and user expectations rising, developers must stay ahead of the curve to create meaningful digital experiences.</p>

<h3>AI Integration in Development</h3>
<p>Artificial Intelligence is no longer just a buzzword—it's becoming an integral part of the development process. From code completion to automated testing, AI is transforming how we build web applications.</p>

<blockquote>
<p>"The future belongs to those who can harness the power of AI while maintaining the human touch in their creations."</p>
</blockquote>

<h3>Edge Computing Revolution</h3>
<p>Edge computing is bringing processing power closer to users, reducing latency and improving performance. This shift is enabling new possibilities for real-time applications and immersive experiences.</p>

<ul>
<li>Faster response times</li>
<li>Reduced server costs</li>
<li>Improved user experience</li>
<li>Enhanced security</li>
</ul>

<h3>WebAssembly: The Game Changer</h3>
<p>WebAssembly is opening new doors for web applications, allowing near-native performance in the browser. This technology is enabling complex applications like video editors, CAD software, and games to run smoothly on the web.</p>

<pre><code>// Example WebAssembly integration
const wasmModule = await WebAssembly.instantiate(wasmBytes);
const result = wasmModule.instance.exports.calculate(values);</code></pre>

<h2>Conclusion</h2>
<p>The future of web development is bright and full of possibilities. By embracing these emerging technologies while keeping user needs at the forefront, we can create web experiences that were once thought impossible.</p>`,
      author: {
        name: "Sarah Johnson",
        bio: "Full-stack developer with 10+ years of experience. Passionate about emerging technologies and their impact on web development. Currently working at TechCorp, helping build the next generation of web applications.",
        avatar: "1497366216548-37526070297c",
        slug: "sarah-johnson"
      },
      publishDate: "2024-01-15",
      readTime: "8 min read",
      category: "Technology",
      tags: ["Web Development", "AI", "Edge Computing", "WebAssembly"],
      image: "https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=1200&h=600&fit=crop&auto=format"
    };

    const mockRelated = [
      {
        id: 2,
        title: "Building Sustainable Digital Products",
        excerpt: "How to create digital products that are not only user-friendly but also environmentally conscious.",
        author: "Mike Chen",
        publishDate: "2024-01-14",
        readTime: "6 min read",
        category: "Design",
        slug: "sustainable-digital-products",
        image: "https://images.unsplash.com/photo-1460925895917-afdab827c52f?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 3,
        title: "The Art of Remote Collaboration",
        excerpt: "Best practices and tools for effective remote team collaboration in the modern distributed workplace.",
        author: "Emily Davis",
        publishDate: "2024-01-13",
        readTime: "5 min read",
        category: "Productivity",
        slug: "remote-collaboration",
        image: "https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=400&h=300&fit=crop&auto=format"
      },
      {
        id: 4,
        title: "AI Integration in Modern Development",
        excerpt: "Exploring how artificial intelligence is revolutionizing the way we build and deploy applications.",
        author: "David Kim",
        publishDate: "2024-01-12",
        readTime: "10 min read",
        category: "Technology",
        slug: "ai-integration-modern-development",
        image: "https://images.unsplash.com/photo-1518770660439-4636190af475?w=400&h=300&fit=crop&auto=format"
      }
    ];

    setTimeout(() => {
      setPost(mockPost);
      setRelatedPosts(mockRelated);
      setLoading(false);
    }, 1000);

    // TODO: connect API endpoint using src/config/api.js
    // fetch(`${API_BASE_URL}${API_ENDPOINTS.POST_BY_SLUG.replace(':slug', slug)}`)
    //   .then(res => res.json())
    //   .then(data => {
    //     setPost(data);
    //     return fetch(`${API_BASE_URL}${API_ENDPOINTS.RELATED_POSTS.replace(':id', data.id)}`);
    //   })
    //   .then(res => res.json())
    //   .then(setRelatedPosts)
    //   .finally(() => setLoading(false));
  }, [slug]);

  const handleShare = (platform) => {
    const url = window.location.href;
    const text = `Check out this article: ${post?.title}`;
    
    switch (platform) {
      case 'twitter':
        window.open(`https://twitter.com/intent/tweet?text=${encodeURIComponent(text)}&url=${encodeURIComponent(url)}`);
        break;
      case 'facebook':
        window.open(`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(url)}`);
        break;
      case 'linkedin':
        window.open(`https://www.linkedin.com/sharing/share-offsite/?url=${encodeURIComponent(url)}`);
        break;
      default:
        navigator.clipboard.writeText(url);
        alert('Link copied to clipboard!');
    }
  };

  const handleCommentSubmit = (e) => {
    e.preventDefault();
    if (!commentText.trim()) return;
    
    const newComment = {
      id: comments.length + 1,
      author: 'You',
      content: commentText,
      timestamp: 'Just now',
      avatar: '1522071820081-009f0129c71c'
    };
    
    setComments([newComment, ...comments]);
    setCommentText('');
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex flex-col">
        <Header />
        <main className="flex-1 flex items-center justify-center">
          <div className="text-center">
            <Icon name="Loader2" className="w-8 h-8 animate-spin text-blue-600 mx-auto mb-4" />
            <p className="text-slate-600">Loading post...</p>
          </div>
        </main>
        <Footer />
      </div>
    );
  }

  if (!post) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex flex-col">
        <Header />
        <main className="flex-1 flex items-center justify-center">
          <div className="text-center">
            <Icon name="FileX" className="w-16 h-16 text-slate-400 mx-auto mb-4" />
            <h2 className="text-2xl font-bold text-slate-800 mb-2">Post Not Found</h2>
            <p className="text-slate-600 mb-4">The post you're looking for doesn't exist.</p>
            <Link to="/blog" className="text-blue-600 hover:text-blue-700 font-medium">
              Back to all posts
            </Link>
          </div>
        </main>
        <Footer />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 flex flex-col">
      <Header />
      
      <main className="flex-1">
        {/* Post Header */}
        <section className="py-12 md:py-24">
          <div className="container mx-auto px-4 md:px-6">
            <div className="max-w-4xl mx-auto">
              <div className="mb-8">
                <Link to="/blog" className="inline-flex items-center gap-2 text-blue-600 hover:text-blue-700 mb-6 transition-colors">
                  <Icon name="ArrowLeft" className="w-5 h-5" />
                  Back to all posts
                </Link>
                
                <div className="flex flex-wrap gap-2 mb-6">
                  {post.tags.map((tag) => (
                    <span key={tag} className="inline-flex items-center gap-1 px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-xs font-medium">
                      <Icon name="Tag" className="w-3 h-3" />
                      {tag}
                    </span>
                  ))}
                </div>

                <h1 className="text-4xl md:text-6xl font-extrabold text-slate-800 mb-6 leading-tight">
                  {post.title}
                </h1>

                <div className="flex items-center gap-6 text-slate-600 mb-8">
                  <div className="flex items-center gap-3">
                    <div className="w-12 h-12 rounded-full overflow-hidden">
                      <img
                        src={`https://images.unsplash.com/photo-${post.author.avatar}?w=128&h=128&fit=crop&auto=format`}
                        alt={post.author.name}
                        className="w-full h-full object-cover"
                        onError={(e) => {
                          e.target.onerror = null;
                          e.target.src = 'https://placehold.co/128x128/1a1a2e/eaeaea?text=Author';
                        }}
                      />
                    </div>
                    <div>
                      <Link to={`/author/${post.author.slug}`} className="font-semibold text-slate-800 hover:text-blue-600 transition-colors">
                        {post.author.name}
                      </Link>
                      <p className="text-sm">
                        {new Date(post.publishDate).toLocaleDateString('en-US', { 
                          month: 'long', 
                          day: 'numeric', 
                          year: 'numeric' 
                        })}
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center gap-4">
                    <span className="flex items-center gap-1">
                      <Icon name="Clock" className="w-4 h-4" />
                      {post.readTime}
                    </span>
                    <span className="flex items-center gap-1">
                      <Icon name="MessageSquare" className="w-4 h-4" />
                      {comments.length} comments
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Post Content */}
        <section className="py-12 bg-white/30 backdrop-blur-sm">
          <div className="container mx-auto px-4 md:px-6">
            <div className="max-w-4xl mx-auto">
              <div 
                className="prose prose-lg max-w-none prose-headings:text-slate-800 prose-p:text-slate-700 prose-a:text-blue-600 prose-strong:text-slate-800 prose-code:text-pink-600 prose-code:bg-slate-100 prose-code:px-1 prose-code:py-0.5 prose-code:rounded prose-blockquote:border-l-blue-500 prose-blockquote:bg-blue-50 prose-blockquote:p-6 prose-blockquote:rounded-xl"
                dangerouslySetInnerHTML={{ __html: post.content }}
              />

              {/* Social Sharing */}
              <div className="mt-12 pt-8 border-t border-slate-200">
                <h3 className="text-lg font-semibold text-slate-800 mb-4">Share this article</h3>
                <div className="flex flex-wrap gap-3">
                  <button
                    onClick={() => handleShare('twitter')}
                    className="inline-flex items-center gap-2 px-4 py-2 bg-white/70 backdrop-blur-lg rounded-xl border border-white/20 text-slate-700 hover:bg-white/90 transition-all duration-300 shadow-soft hover:shadow-lg"
                  >
                    <Icon name="Twitter" className="w-5 h-5" />
                    Twitter
                  </button>
                  <button
                    onClick={() => handleShare('facebook')}
                    className="inline-flex items-center gap-2 px-4 py-2 bg-white/70 backdrop-blur-lg rounded-xl border border-white/20 text-slate-700 hover:bg-white/90 transition-all duration-300 shadow-soft hover:shadow-lg"
                  >
                    <Icon name="Facebook" className="w-5 h-5" />
                    Facebook
                  </button>
                  <button
                    onClick={() => handleShare('linkedin')}
                    className="inline-flex items-center gap-2 px-4 py-2 bg-white/70 backdrop-blur-lg rounded-xl border border-white/20 text-slate-700 hover:bg-white/90 transition-all duration-300 shadow-soft hover:shadow-lg"
                  >
                    <Icon name="Linkedin" className="w-5 h-5" />
                    LinkedIn
                  </button>
                  <button
                    onClick={() => handleShare('copy')}
                    className="inline-flex items-center gap-2 px-4 py-2 bg-white/70 backdrop-blur-lg rounded-xl border border-white/20 text-slate-700 hover:bg-white/90 transition-all duration-300 shadow-soft hover:shadow-lg"
                  >
                    <Icon name="Copy" className="w-5 h-5" />
                    Copy Link
                  </button>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Author Bio */}
        <section className="py-12">
          <div className="container mx-auto px-4 md:px-6">
            <div className="max-w-4xl mx-auto">
              <div className="bg-white/70 backdrop-blur-lg rounded-xl shadow-lg border border-white/20 p-8">
                <div className="flex items-start gap-6">
                  <div className="w-20 h-20 rounded-full overflow-hidden flex-shrink-0">
                    <img
                      src={`https://images.unsplash.com/photo-${post.author.avatar}?w=160&h=160&fit=crop&auto=format`}
                      alt={post.author.name}
                      className="w-full h-full object-cover"
                      onError={(e) => {
                        e.target.onerror = null;
                        e.target.src = 'https://placehold.co/160x160/1a1a2e/eaeaea?text=Author';
                      }}
                    />
                  </div>
                  <div className="flex-1">
                    <h3 className="text-xl font-bold text-slate-800 mb-2">{post.author.name}</h3>
                    <p className="text-slate-600 mb-4 leading-relaxed">{post.author.bio}</p>
                    <Link
                      to={`/author/${post.author.slug}`}
                      className="inline-flex items-center gap-2 text-blue-600 hover:text-blue-700 font-medium transition-colors"
                    >
                      View Profile
                      <Icon name="ArrowRight" className="w-4 h-4" />
                    </Link>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Related Posts */}
        <section className="py-12 bg-white/30 backdrop-blur-sm">
          <div className="container mx-auto px-4 md:px-6">
            <div className="max-w-6xl mx-auto">
              <h2 className="text-3xl font-bold text-slate-800 mb-8">Related Articles</h2>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6 md:gap-8">
                {relatedPosts.map((relatedPost) => (
                  <article key={relatedPost.id} className="bg-white/70 backdrop-blur-lg rounded-xl shadow-lg border border-white/20 overflow-hidden hover:shadow-xl transition-all duration-300 transform hover:scale-[1.02]">
                    <div className="aspect-video relative overflow-hidden">
                      <img
                        src={relatedPost.image}
                        alt={relatedPost.title}
                        className="w-full h-full object-cover"
                        onError={(e) => {
                          e.target.onerror = null;
                          e.target.src = 'https://placehold.co/400x300/1a1a2e/eaeaea?text=Image';
                        }}
                      />
                      <div className="absolute top-4 left-4">
                        <span className="inline-flex items-center gap-1 px-3 py-1 bg-white/80 backdrop-blur-sm rounded-full text-xs font-medium text-slate-700 border border-white/50">
                          <Icon name="Tag" className="w-3 h-3" />
                          {relatedPost.category}
                        </span>
                      </div>
                    </div>
                    <div className="p-6">
                      <h3 className="text-xl font-bold text-slate-800 mb-2">
                        <Link to={`/blog/${relatedPost.slug}`} className="hover:text-blue-600 transition-colors">
                          {relatedPost.title}
                        </Link>
                      </h3>
                      <p className="text-slate-600 mb-4 line-clamp-3">{relatedPost.excerpt}</p>
                      <div className="flex items-center justify-between text-sm text-slate-500">
                        <div className="flex items-center gap-4">
                          <span className="flex items-center gap-1">
                            <Icon name="User" className="w-4 h-4" />
                            {relatedPost.author}
                          </span>
                          <span className="flex items-center gap-1">
                            <Icon name="Clock" className="w-4 h-4" />
                            {relatedPost.readTime}
                          </span>
                        </div>
                        <time>{new Date(relatedPost.publishDate).toLocaleDateString()}</time>
                      </div>
                    </div>
                  </article>
                ))}
              </div>
            </div>
          </div>
        </section>

        {/* Comments Section */}
        <section className="py-12">
          <div className="container mx-auto px-4 md:px-6">
            <div className="max-w-4xl mx-auto">
              <h2 className="text-3xl font-bold text-slate-800 mb-8">Comments ({comments.length})</h2>
              
              {/* Comment Form */}
              <div className="bg-white/70 backdrop-blur-lg rounded-xl shadow-lg border border-white/20 p-6 mb-8">
                <h3 className="text-lg font-semibold text-slate-800 mb-4">Leave a Comment</h3>
                <form onSubmit={handleCommentSubmit} className="space-y-4">
                  <textarea
                    value={commentText}
                    onChange={(e) => setCommentText(e.target.value)}
                    placeholder="Share your thoughts..."
                    className="w-full px-4 py-3 rounded-lg border border-slate-200 bg-white/50 backdrop-blur-sm focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 resize-none"
                    rows="4"
                  />
                  <button
                    type="submit"
                    disabled={!commentText.trim()}
                    className="px-6 py-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-lg font-medium hover:from-blue-700 hover:to-purple-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300"
                  >
                    Post Comment
                  </button>
                </form>
              </div>

              {/* Comments List */}
              <div className="space-y-6">
                {comments.map((comment) => (
                  <div key={comment.id} className="bg-white/70 backdrop-blur-lg rounded-xl shadow-lg border border-white/20 p-6">
                    <div className="flex items-start gap-4">
                      <div className="w-10 h-10 rounded-full overflow-hidden flex-shrink-0">
                        <img
                          src={`https://images.unsplash.com/photo-${comment.avatar}?w=80&h=80&fit=crop&auto=format`}
                          alt={comment.author}
                          className="w-full h-full object-cover"
                          onError={(e) => {
                            e.target.onerror = null;
                            e.target.src = 'https://placehold.co/80x80/1a1a2e/eaeaea?text=User';
                          }}
                        />
                      </div>
                      <div className="flex-1">
                        <div className="flex items-center justify-between mb-2">
                          <h4 className="font-semibold text-slate-800">{comment.author}</h4>
                          <span className="text-sm text-slate-500">{comment.timestamp}</span>
                        </div>
                        <p className="text-slate-700">{comment.content}</p>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </section>
      </main>

      <Footer />
    </div>
  );
}