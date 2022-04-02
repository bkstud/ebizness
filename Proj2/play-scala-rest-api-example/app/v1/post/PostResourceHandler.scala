package v1.post

import javax.inject.{Inject, Provider}

import play.api.MarkerContext

import scala.concurrent.{ExecutionContext, Future}
import play.api.libs.json._

/**
  * DTO for displaying post information.
  */
case class PostResource(id: String, link: String, title: String, body: String)

object PostResource {
  /**
    * Mapping to read/write a PostResource out as a JSON value.
    */
    implicit val format: Format[PostResource] = Json.format
}


/**
  * Controls access to the backend data, returning [[PostResource]]
  */
class PostResourceHandler @Inject()(
    routerProvider: Provider[PostRouter],
    postRepository: PostRepository)(implicit ec: ExecutionContext) {

  def create(postInput: PostFormInput)(
      implicit mc: MarkerContext): Future[PostResource] = {
    val Ids = postRepository.postList.map {el => el.id.underlying}
    // generated new unique id for post
    val data = PostData(PostId(s"${Ids.max + 1}"), postInput.title, postInput.body)
    postRepository.create(data).map { id =>
      createPostResource(data)
    }
  }

  def lookup(id: String)(
      implicit mc: MarkerContext): Future[Option[PostResource]] = {
    val postFuture = postRepository.get(PostId(id))
    postFuture.map { maybePostData =>
      maybePostData.map { postData =>
        createPostResource(postData)
      }
    }
  }

  def delete(id: String)(
      implicit mc: MarkerContext): Future[Option[PostResource]] = {
    val postFuture = postRepository.delete(PostId(id))
    postFuture.map { maybePostData =>
      maybePostData.map { postData =>
        createPostResource(postData)
      }
    }
  }

  def update(id: String, postInput: PostFormInput)(
      implicit mc: MarkerContext): Future[Option[PostResource]] = {
    val data = PostData(PostId(id), postInput.title, postInput.body)
    postRepository.update(PostId(id), data).map { maybePostData =>
      maybePostData.map { postData =>
        createPostResource(postData)
      }
    }
  }

  def find(implicit mc: MarkerContext): Future[Iterable[PostResource]] = {
    postRepository.list().map { postDataList =>
      postDataList.map(postData => createPostResource(postData))
    }
  }

  private def createPostResource(p: PostData): PostResource = {
    PostResource(p.id.toString, routerProvider.get.link(p.id), p.title, p.body)
  }

}
