# Event의 Thread Safety

```csharp
public class Publisher
{
    private volatile EventHandler<MyEventArgs> _myEvent;
    
    public event EventHandler<MyEventArgs> MyEvent
    {
        add
        {
            EventHandler<MyEventArgs> oldHandler;
            EventHandler<MyEventArgs> newHandler;
            do
            {
                oldHandler = _myEvent;
                newHandler = (EventHandler<MyEventArgs>)Delegate.Combine(oldHandler, value);
            }
            while (Interlocked.CompareExchange(ref _myEvent, newHandler, oldHandler) != oldHandler);
        }
        remove
        {
            EventHandler<MyEventArgs> oldHandler;
            EventHandler<MyEventArgs> newHandler;
            do
            {
                oldHandler = _myEvent;
                newHandler = (EventHandler<MyEventArgs>)Delegate.Remove(oldHandler, value);
            }
            while (Interlocked.CompareExchange(ref _myEvent, newHandler, oldHandler) != oldHandler);
        }
    }
}
```