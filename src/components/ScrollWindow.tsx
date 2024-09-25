import './css/ScrollWindow.css'; // Import the CSS for the scrollable component
import CategoryScrollItem from './CategoryScrollItem';

const ScrollWindow = () => {
    const sample_data = {
        categories: [
            {
                categoryName: 'Technology',
                posts: [
                    {
                        id: 1,
                        title: 'The Rise of AI',
                        content:
                            'An in-depth look at how AI is shaping the future.',
                    },
                    {
                        id: 2,
                        title: 'Blockchain Basics',
                        content:
                            'Understanding the fundamentals of blockchain technology.',
                    },
                    {
                        id: 3,
                        title: 'Quantum Computing',
                        content:
                            'Exploring the potential of quantum computers.',
                    },
                    {
                        id: 4,
                        title: 'Cybersecurity Trends',
                        content:
                            'Top trends in cybersecurity you should know about.',
                    },
                    {
                        id: 5,
                        title: 'The Internet of Things',
                        content:
                            'How IoT is connecting the world.',
                    },
                ],
            },
            {
                categoryName: 'Science',
                posts: [
                    {
                        id: 6,
                        title: 'The Future of Space Exploration',
                        content:
                            "What's next for space exploration?",
                    },
                    {
                        id: 7,
                        title: 'CRISPR and Gene Editing',
                        content:
                            'How CRISPR is revolutionizing genetic science.',
                    },
                    {
                        id: 8,
                        title: 'Climate Change and Its Effects',
                        content:
                            'Understanding the global impact of climate change.',
                    },
                    {
                        id: 9,
                        title: 'Renewable Energy Advances',
                        content:
                            'Latest innovations in renewable energy sources.',
                    },
                    {
                        id: 10,
                        title: 'The Mysteries of Dark Matter',
                        content:
                            'Unraveling the secrets of dark matter.',
                    },
                ],
            },
        ],
    };

    // Need to parse the JSON

    return (
        <div className="scrollable-sidebar">
            {/* Add content to the scrollable sidebar */}
            {sample_data.categories.map(
                (category, index) => (
                    <CategoryScrollItem
                        key={index}
                        categoryName={
                            category.categoryName
                        }
                        posts={category.posts}
                    />
                )
            )}

        </div>
    );
};

export default ScrollWindow;
