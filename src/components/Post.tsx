import './css/Post.css';

type PostProps = {
  post: {
    ID: number;
    Title: string;
    Url: string;
    Body: string;
    Date: string;
  };
};

function Post({ post }: PostProps) {
  return (
    <div className="post-container">
      <h2 className="post-title">{post.Title}</h2>
      <p className="post-date">Published on: {new Date(post.Date).toLocaleDateString()}</p>
      <p className="post-body">{post.Body}</p>
      <a href={post.Url} className="post-link" target="_blank" rel="noopener noreferrer">
        Read More
      </a>
    </div>
  );
}

export default Post;

