import { FC } from 'react';
import Link from 'next/link';

interface ProjectProps {
  id: number;
  name: string;
  description: string;
  tags?: string[];
  stars?: number;
  forks?: number;
  updatedAt?: string;
}

export const ProjectCard: FC<ProjectProps> = ({
  id,
  name,
  description,
  tags,
  stars = 0,
  forks = 0,
  updatedAt
}) => {
  return (
    <div className="rounded-xl bg-white dark:bg-gray-700 shadow-md p-5 hover:shadow-lg transition-shadow">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        {/* 左侧：项目信息 */}
        <div className="flex-1">
          <div className="flex items-center gap-2 mb-2">
            <Link href={`/${id}`}>
              <h2 className="text-xl font-semibold text-gray-800 dark:text-gray-200 hover:text-blue-600 dark:hover:text-blue-400">
                {name}
              </h2>
            </Link>
          </div>
          <p className="text-gray-600 dark:text-gray-300 mb-3 line-clamp-2">
            {description}
          </p>
          
          {/* 标签列表 */}
          <div className="flex flex-wrap gap-2">
            {tags?.map((tag, index) => (
              <span
                key={index}
                className="inline-block bg-gray-100 dark:bg-gray-700 rounded-full px-3 py-1 text-sm font-medium text-gray-700 dark:text-gray-300"
              >
                {tag}
              </span>
            ))}
          </div>
        </div>

        {/* 右侧：统计信息 */}
        <div className="flex flex-row md:flex-col items-center gap-4 md:gap-2">
          <div className="flex items-center gap-4">
            <div className="flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-gray-500 dark:text-gray-400 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
              </svg>
              <span className="text-gray-700 dark:text-gray-300">{stars}</span>
            </div>
            <div className="flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-gray-500 dark:text-gray-400 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
              </svg>
              <span className="text-gray-700 dark:text-gray-300">{forks}</span>
            </div>
          </div>
          <div className="text-sm text-gray-500 dark:text-gray-400">
            最后更新于 {updatedAt}
          </div>
        </div>
      </div>
    </div>
  );
};