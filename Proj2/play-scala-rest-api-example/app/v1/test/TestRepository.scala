package v1.test

import javax.inject.{Inject, Singleton}

import akka.actor.ActorSystem
import play.api.libs.concurrent.CustomExecutionContext
import play.api.{Logger, MarkerContext}

import scala.concurrent.Future
import scala.collection.mutable.ListBuffer

final case class TestData(id: TestId, title: String, body: String)

class TestId private (val underlying: Int) extends AnyVal {
  override def toString: String = underlying.toString
}

object TestId {
  def apply(raw: String): TestId = {
    require(raw != null)
    new TestId(Integer.parseInt(raw))
  }
}

class TestExecutionContext @Inject()(actorSystem: ActorSystem)
    extends CustomExecutionContext(actorSystem, "repository.dispatcher")

/**
  * A pure non-blocking interface for the TestRepository.
  */
trait TestRepository {
  def create(data: TestData)(implicit mc: MarkerContext): Future[TestId]

  // show all
  def list()(implicit mc: MarkerContext): Future[Iterable[TestData]]

  def get(id: TestId)(implicit mc: MarkerContext): Future[Option[TestData]]
  def delete(id: TestId)(implicit mc: MarkerContext): Future[Option[TestData]]
  def update(id: TestId, data: TestData)(implicit mc: MarkerContext): Future[Option[TestData]]
  
  val testList: ListBuffer[TestData]
}

/**
  * A trivial implementation for the Test Repository.
  *
  * A custom execution context is used here to establish that blocking operations should be
  * executed in a different thread than Play's ExecutionContext, which is used for CPU bound tasks
  * such as rendering.
  */
@Singleton
class TestRepositoryImpl @Inject()()(implicit ec: TestExecutionContext)
    extends TestRepository {

  private val logger = Logger(this.getClass)

  val testList = ListBuffer(
    TestData(TestId("1"), "title 1", "blog test 1"),
    TestData(TestId("2"), "title 2", "blog test 2"),
    TestData(TestId("3"), "title 3", "blog test 3"),
    TestData(TestId("4"), "title 4", "blog test 4"),
    TestData(TestId("5"), "title 5", "blog test 5")
  )

  override def list()(
      implicit mc: MarkerContext): Future[Iterable[TestData]] = {
    Future {
      logger.trace(s"list: ")
      testList
    }
  }

  override def get(id: TestId)(
      implicit mc: MarkerContext): Future[Option[TestData]] = {
    Future {
      logger.trace(s"get: id = $id")
      testList.find(test => test.id == id)
    }
  }

  def create(data: TestData)(implicit mc: MarkerContext): Future[TestId] = {
    Future {
      logger.trace(s"create: data = $data")
      testList += data
      data.id
    }
  }

  def delete(id: TestId)(implicit mc: MarkerContext): Future[Option[TestData]] = {
    Future {
      logger.trace(s"delete: id = $id")
      
      var index = testList.indexWhere(test => test.id == id)
      var out: TestData = null
      if(index != -1) {
       out = testList.remove(index)
      }
      Option(out)
    }

  }

  override def update(id: TestId, data: TestData)(
      implicit mc: MarkerContext): Future[Option[TestData]] = {
    Future {
      logger.trace(s"update: id = $id")
      logger.trace(s"update: data = $data")
      
      var index = testList.indexWhere(test => test.id == id)
      var out: TestData = null
      if(index != -1) {
       
       out = TestData(testList.remove(index).id, data.title, data.body)
       testList += out
      }
      Option(out)
    }
  }

}
