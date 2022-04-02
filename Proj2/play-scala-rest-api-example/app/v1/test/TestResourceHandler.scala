package v1.test

import javax.inject.{Inject, Provider}

import play.api.MarkerContext

import scala.concurrent.{ExecutionContext, Future}
import play.api.libs.json._

/**
  * DTO for displaying test information.
  */
case class TestResource(id: String, link: String, title: String, body: String)

object TestResource {
  /**
    * Mapping to read/write a TestResource out as a JSON value.
    */
    implicit val format: Format[TestResource] = Json.format
}


/**
  * Controls access to the backend data, returning [[TestResource]]
  */
class TestResourceHandler @Inject()(
    routerProvider: Provider[TestRouter],
    testRepository: TestRepository)(implicit ec: ExecutionContext) {

  def create(testInput: TestFormInput)(
      implicit mc: MarkerContext): Future[TestResource] = {
    val Ids = testRepository.testList.map {el => el.id.underlying}
    // generated new unique id for test
    val data = TestData(TestId(s"${Ids.max + 1}"), testInput.title, testInput.body)
    testRepository.create(data).map { id =>
      createTestResource(data)
    }
  }

  def lookup(id: String)(
      implicit mc: MarkerContext): Future[Option[TestResource]] = {
    val testFuture = testRepository.get(TestId(id))
    testFuture.map { maybeTestData =>
      maybeTestData.map { testData =>
        createTestResource(testData)
      }
    }
  }

  def delete(id: String)(
      implicit mc: MarkerContext): Future[Option[TestResource]] = {
    val testFuture = testRepository.delete(TestId(id))
    testFuture.map { maybeTestData =>
      maybeTestData.map { testData =>
        createTestResource(testData)
      }
    }
  }

  def update(id: String, testInput: TestFormInput)(
      implicit mc: MarkerContext): Future[Option[TestResource]] = {
    val data = TestData(TestId(id), testInput.title, testInput.body)
    testRepository.update(TestId(id), data).map { maybeTestData =>
      maybeTestData.map { testData =>
        createTestResource(testData)
      }
    }
  }

  def find(implicit mc: MarkerContext): Future[Iterable[TestResource]] = {
    testRepository.list().map { testDataList =>
      testDataList.map(testData => createTestResource(testData))
    }
  }

  private def createTestResource(p: TestData): TestResource = {
    TestResource(p.id.toString, routerProvider.get.link(p.id), p.title, p.body)
  }

}
