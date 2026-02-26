import React from 'react';
import { Link } from 'react-router-dom';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function Footer() {
  return (
    <footer className="bg-white/50 backdrop-blur-xl border-t border-blue-100/30 mt-24">
      <div className="container mx-auto px-4 md:px-6 py-12">
        <div className="grid md:grid-cols-4 gap-8">
          <div className="md:col-span-2">
            <div className="flex items-center gap-3 mb-4">
              <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-blue-500 to-purple-500 flex items-center justify-center shadow-soft">
                <Icon name="PenTool" className="w-5 h-5 text-white" />
              </div>
              <span className="text-xl font-bold text-slate-800">BlogSphere</span>
            </div>
            <p className="text-slate-600 max-w-md leading-relaxed">
              A modern platform for creators to share stories, insights, and ideas with the world.
            </p>
          </div>
          <div>
            <h4 className="font-semibold text-slate-800 mb-3">Quick Links</h4>
            <ul className="space-y-2 text-slate-600">
              <li><Link to="/" className="hover:text-blue-600 transition-colors">Home</Link></li>
              <li><Link to="/blog" className="hover:text-blue-600 transition-colors">Blog</Link></li>
              <li><Link to="/login" className="hover:text-blue-600 transition-colors">Login</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold text-slate-800 mb-3">Connect</h4>
            <div className="flex gap-3">
              <a href="#" className="w-10 h-10 rounded-lg bg-white/70 backdrop-blur-sm border border-blue-100/30 flex items-center justify-center text-slate-600 hover:text-blue-600 hover:border-blue-300 transition-all duration-300 shadow-soft hover:shadow-lg">
                <Icon name="Twitter" className="w-5 h-5" />
              </a>
              <a href="#" className="w-10 h-10 rounded-lg bg-white/70 backdrop-blur-sm border border-blue-100/30 flex items-center justify-center text-slate-600 hover:text-blue-600 hover:border-blue-300 transition-all duration-300 shadow-soft hover:shadow-lg">
                <Icon name="Github" className="w-5 h-5" />
              </a>
              <a href="#" className="w-10 h-10 rounded-lg bg-white/70 backdrop-blur-sm border border-blue-100/30 flex items-center justify-center text-slate-600 hover:text-blue-600 hover:border-blue-300 transition-all duration-300 shadow-soft hover:shadow-lg">
                <Icon name="Linkedin" className="w-5 h-5" />
              </a>
            </div>
          </div>
        </div>
        <div className="border-t border-blue-100/30 mt-12 pt-8 text-center text-slate-500">
          © 2025 BlogSphere. All rights reserved.
        </div>
      </div>
    </footer>
  );
}