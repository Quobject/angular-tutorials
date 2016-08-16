using Xunit;

//./packages/xunit.runner.console.2.1.0/tools/xunit.console.exe ./bin/Debug/MyFirstUnitTests.dll


namespace MyFirstUnitTests
{
    public class Class1
    {
        [Fact]
        public void PassingTest()
        {
            Assert.Equal(4, Add(2, 2));
        }

        [Fact]
        public void FailingTest()
        {
            Assert.Equal(5, Add(2, 2));
        }

        [Fact]
        public void ProgramTest()
        {
            var value = ConsoleApplication1.Program.Add(3);
            Assert.Equal(4,value);
        }

        int Add(int x, int y)
        {
            return x + y;
        }

        [Theory]
        [InlineData(3)]
        [InlineData(5)]
        [InlineData(7)]
        public void MyFirstTheory(int value)
        {
            Assert.True(IsOdd(value));
        }

        bool IsOdd(int value)
        {
            return value % 2 == 1;
        }

        

    }



}