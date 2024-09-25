import './css/ScrollWindow.css'; // Import the CSS for the scrollable component
import CategoryScrollItem from './CategoryScrollItem';
import sample_api_data from '../assets/sample_api_data';
type ApiData = {
    ID: number;
    Title: string;
    Url: string;
    Body: string;
    Date: string;
    Category: string;
};

type Post = {
    ID: number;
    Title: string;
    Url: string;
    Body: string;
    Date: string;
};
const ScrollWindow = () => {
    // required data format
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
    const transformData = (data: ApiData[]) => {
    // Create a map to store categories and their posts
    const categoryMap: {
        [key: string]: {
            categoryName: string;
            posts: Post[];
        };
    } = {};

    data.forEach((item) => {
        const { ID, Title, Body, Url, Category } = item;

        // If the category doesn't exist in the map, create a new entry
        if (!categoryMap[Category]) {
            categoryMap[Category] = {
                categoryName: Category,
                posts: [],
            };
        }

        // Push the post into the correct category's posts array
        categoryMap[Category].posts.push({
            ID: ID,
            Title: Title,
            Body: Body,
            Url: Url,
            Date: item.Date,  // Including the date field from ApiData
        });
    });

    // Convert the map to an array
    return {
        categories: Object.values(categoryMap),
    };
};
    const transformedData = transformData(
        sample_api_data
    );

    console.log(transformedData)

    return (
        <div className="scrollable-sidebar">
            {/* Add content to the scrollable sidebar */}
            {transformedData.categories.map(
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
