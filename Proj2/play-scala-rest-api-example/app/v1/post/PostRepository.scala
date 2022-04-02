package v1.post

import javax.inject.{Inject, Singleton}

import akka.actor.ActorSystem
import play.api.libs.concurrent.CustomExecutionContext
import play.api.{Logger, MarkerContext}

import scala.concurrent.Future
import scala.collection.mutable.ListBuffer

final case class PostData(id: PostId, title: String, body: String)

class PostId private (val underlying: Int) extends AnyVal {
  override def toString: String = underlying.toString
}

object PostId {
  def apply(raw: String): PostId = {
    require(raw != null)
    new PostId(Integer.parseInt(raw))
  }
}

class PostExecutionContext @Inject()(actorSystem: ActorSystem)
    extends CustomExecutionContext(actorSystem, "repository.dispatcher")

/**
  * A pure non-blocking interface for the PostRepository.
  */
trait PostRepository {
  def create(data: PostData)(implicit mc: MarkerContext): Future[PostId]

  // show all
  def list()(implicit mc: MarkerContext): Future[Iterable[PostData]]

  def get(id: PostId)(implicit mc: MarkerContext): Future[Option[PostData]]
  def delete(id: PostId)(implicit mc: MarkerContext): Future[Option[PostData]]
  def update(id: PostId, data: PostData)(implicit mc: MarkerContext): Future[Option[PostData]]
  
  val postList: ListBuffer[PostData]
}

/**
  * A trivial implementation for the Post Repository.
  *
  * A custom execution context is used here to establish that blocking operations should be
  * executed in a different thread than Play's ExecutionContext, which is used for CPU bound tasks
  * such as rendering.
  */
@Singleton
class PostRepositoryImpl @Inject()()(implicit ec: PostExecutionContext)
    extends PostRepository {

  private val logger = Logger(this.getClass)

  val postList = ListBuffer(
    PostData(PostId("1"), "title 1", "blog post 1"),
    PostData(PostId("2"), "title 2", "blog post 2"),
    PostData(PostId("3"), "title 3", "blog post 3"),
    PostData(PostId("4"), "title 4", "blog post 4"),
    PostData(PostId("5"), "title 5", "blog post 5")
  )

  override def list()(
      implicit mc: MarkerContext): Future[Iterable[PostData]] = {
    Future {
      logger.trace(s"list: ")
      postList
    }
  }

  override def get(id: PostId)(
      implicit mc: MarkerContext): Future[Option[PostData]] = {
    Future {
      logger.trace(s"get: id = $id")
      postList.find(post => post.id == id)
    }
  }

  def create(data: PostData)(implicit mc: MarkerContext): Future[PostId] = {
    Future {
      logger.trace(s"create: data = $data")
      postList += data
      data.id
    }
  }

  def delete(id: PostId)(implicit mc: MarkerContext): Future[Option[PostData]] = {
    Future {
      logger.trace(s"delete: id = $id")
      
      var index = postList.indexWhere(post => post.id == id)
      var out: PostData = null
      if(index != -1) {
       out = postList.remove(index)
      }
      Option(out)
    }

  }

  override def update(id: PostId, data: PostData)(
      implicit mc: MarkerContext): Future[Option[PostData]] = {
    Future {
      logger.trace(s"update: id = $id")
      logger.trace(s"update: data = $data")
      
      var index = postList.indexWhere(post => post.id == id)
      var out: PostData = null
      if(index != -1) {
       
       out = PostData(postList.remove(index).id, data.title, data.body)
       postList += out
      }
      Option(out)
    }
  }

}
