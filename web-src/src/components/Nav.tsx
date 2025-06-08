"use client";

import React, { useState } from "react";
import {
    Settings,
    LayoutDashboard,
    Link2,
    Server,
    Activity,
    Globe,
    Github,
    Snowflake,
    Monitor,
    Triangle,
    Users,
    Plus,
    ChevronDown,
} from "lucide-react";
import DropdownMenu, { DropdownMenuItem } from "@/components/reusable/DropdownMenu";

const teams = [
    {
        id: "team1",
        name: "snowykami's projects",
        icon: <span className="inline-block w-4 h-4 rounded-full bg-gradient-to-tr from-green-400 to-green-600 mr-2" />,
        projects: [
            { id: "blog", name: "blog", icon: <Snowflake className="w-4 h-4 mr-2 text-blue-400" /> },
            { id: "github-readme-stats", name: "github-readme-stats", icon: <Github className="w-4 h-4 mr-2" /> },
            { id: "uptime-kuma", name: "uptime-kuma", icon: <Monitor className="w-4 h-4 mr-2 text-green-400" /> },
            { id: "v0-react-open-source-homepage", name: "v0-react-open-source-homepage", icon: <Triangle className="w-4 h-4 mr-2" /> },
        ],
    },
];

const navItems = [
    { name: "Overview", href: "#", icon: <LayoutDashboard className="w-4 h-4 mr-1" /> },
    { name: "Integrations", href: "#", icon: <Link2 className="w-4 h-4 mr-1" /> },
    { name: "Deployments", href: "#", icon: <Server className="w-4 h-4 mr-1" /> },
    { name: "Activity", href: "#", icon: <Activity className="w-4 h-4 mr-1" /> },
    { name: "Domains", href: "#", icon: <Globe className="w-4 h-4 mr-1" /> },
];

// 构造菜单树
const menuTree: DropdownMenuItem[] = [
    {
        key: "team1",
        label: (
            <span className="flex items-center">
                <span className="inline-block w-4 h-4 rounded-full bg-gradient-to-tr from-green-400 to-green-600 mr-2" />
                <span className="ml-2 px-2 py-0.5 text-xs rounded bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400">Hobby</span>
            </span>
        ),
        icon: <Users className="w-4 h-4" />,
        children: [
            ...teams[0].projects.map((p) => ({
                key: p.id,
                label: (
                    <span className="flex items-center">
                        {p.icon}
                        {p.name}
                    </span>
                ),
                icon: p.icon,
            })),
            {
                key: "create-project",
                label: (
                    <span className="flex items-center text-blue-600 dark:text-blue-400">
                        <Plus className="w-4 h-4 mr-1" /> Create Project
                    </span>
                ),
                icon: <Plus className="w-4 h-4" />,
            },
        ],
    },
    {
        key: "create-team",
        label: (
            <span className="flex items-center text-blue-600 dark:text-blue-400">
                <Plus className="w-4 h-4 mr-1" /> Create Team
            </span>
        ),
        icon: <Plus className="w-4 h-4" />,
    },
];

export default function Nav() {
    const [selectedTeam] = useState(teams[0]);
    const [selectedProject, setSelectedProject] = useState(teams[0].projects[0]);

    return (
        <nav className="pt-4 px-4 w-full border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-900">
            {/* 顶部栏 */}
            <div className="flex items-center justify-between w-full px-4 py-2">
                {/* 左侧：团队/项目下拉菜单 */}
                <DropdownMenu
                    trigger={
                        <button className="flex items-center px-2 py-1 rounded hover:bg-gray-100 dark:hover:bg-gray-800 transition">
                            <span className="mr-2">{selectedTeam.icon}</span>
                            <span className="font-semibold text-gray-900 dark:text-white">{selectedTeam.name}</span>
                            <span className="ml-2 px-2 py-0.5 text-xs rounded bg-gray-100 dark:bg-gray-800 text-gray-500 dark:text-gray-400">Hobby</span>
                            <ChevronDown className="w-4 h-4 mx-2 text-gray-400" />
                            <span className="text-gray-400">/</span>
                            <span className="flex items-center ml-2">
                                {selectedProject.icon}
                                <span className="text-gray-900 dark:text-white">{selectedProject.name}</span>
                            </span>
                            <ChevronDown className="w-4 h-4 ml-1 text-gray-400" />
                        </button>
                    }
                    menu={menuTree}
                    onSelect={(item) => {
                        // 选中项目
                        const project = teams[0].projects.find((p) => p.id === item.key);
                        if (project) setSelectedProject(project);
                    }}
                />
                {/* 右侧设置按钮 */}
                <div className="flex items-center">
                    <button className="flex items-center text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white">
                        <Settings className="w-4 h-4 mr-1" />
                        Settings
                    </button>
                </div>
            </div>
            {/* 下部菜单 */}
            <div>
                <ul className="flex space-x-4 px-4 py-2">
                    {navItems.map((item) => (
                        <li key={item.name}>
                            <a
                                href={item.href}
                                className="flex items-center text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white"
                            >
                                {item.icon}
                                {item.name}
                            </a>
                        </li>
                    ))}
                </ul>
            </div>
        </nav>
    );
}