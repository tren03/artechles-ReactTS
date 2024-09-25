import { useState } from 'react';
import './css/categoryScrollItem.css';
import CategoryScrollSubItem from './CategoryScrollSubItem';
type Post = {
    id: number;
    title: string;
    content: string;
};

type CategoryItem = {
    categoryName: string;
    posts: Post[];
};

function CategoryScrollItem({
    categoryName,
    posts,
}: CategoryItem) {
    const [isDropDown, setDropDown] =
        useState(false);
    // Event handler for button click

    let SubItems: JSX.Element | null;

    if (isDropDown) {
        SubItems = (
            <div>
                {posts.map((post) => {
                    return (
                        <CategoryScrollSubItem
                            key={post.id}
                            title={post.title}
                        />
                    );
                })}
            </div>
        );
    }
    function handleOnClickCategory() {
        console.log('clicked');
        setDropDown(!isDropDown);
    }

    return (
        <div>
            <button
                className="CategoryItem"
                onClick={handleOnClickCategory}
            >
                {categoryName}
            </button>
            {SubItems}
        </div>
    );
}

export default CategoryScrollItem;
