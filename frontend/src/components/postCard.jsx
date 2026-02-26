import React from 'react';
import { Link } from 'react-router-dom';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function PostCard({ post, featured = false }) {
  return (
    <article className={`group bg-white/60 backdrop-blur-xl border border-blue-100/30 rounded-xl overflow-hidden shadow-soft hover:shadow-xl transition-all duration-300 hover:-translate-y-1 ${featured ? 'lg:flex lg:gap-6' : ''}`}>
      <div className={`relative ${featured ? 'lg:w-1/2' : 'aspect-video'}`}>
        <img
          src={post.image}
          alt={post.title}
          className="w-full h-full object-cover"
          onError={(e) => {
            e.target.onerror = null;
            e.target.src = 'https://placehold.co/800x600/f1f5f9/64748b?text=Post';
          }}
        />
        <div className="absolute top-4 left-4">
          <span className="inline-flex items-center gap-1 px-3 py-1 bg-white/80 backdrop-blur-sm rounded-full text-xs font-medium text-slate-700 border border-white/50">
            <Icon name="Tag" className="w-3 h-3" />
            {post.category}
          </span>
        </div>
      </div>
      <div className={`p-6 ${featured ? 'lg:w-1/2 lg:p-8' : ''}`}>
        <div className="flex items-center gap-4 text-sm text-slate-500 mb-3">
          <div className="flex items-center gap-1">
            <Icon name="User" className="w-4 h-4" />
            <span>{post.author}</span>
          </div>
          <div className="flex items-center gap-1">
            <Icon name="Clock" className="w-4 h-4" />
            <span>{post.readTime}</span>
          </div>
        </div>
        <h3 className={`font-bold text-slate-800 mb-3 ${featured ? 'text-2xl lg:text-3xl' : 'text-xl'} group-hover:text-blue-600 transition-colors`}>
          <Link to={`/blog/${post.slug}`}>{post.title}</Link>
        </h3>
        <p className="text-slate-600 leading-relaxed mb-4 line-clamp-3">{post.excerpt}</p>
        <div className="flex items-center justify-between">
          <time className="text-sm text-slate-500">{post.date}</time>
          <Link
            to={`/blog/${post.slug}`}
            className="inline-flex items-center gap-2 text-blue-600 font-medium hover:text-blue-700 transition-colors"
          >
            Read more
            <Icon name="ArrowRight" className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
          </Link>
        </div>
      </div>
    </article>
  );
}