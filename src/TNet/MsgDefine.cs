using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using Knight.Hotfix.Core;
using Knight.Framework.Net;
namespace Game.Net
{
public class CommonBool : IHotfixMessage
{
public bool Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeBOOL(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeBOOL(rBuffer, ref nOffset);

}
}
public class CommonN8 : IHotfixMessage
{
public byte Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN8(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeN8(rBuffer, ref nOffset);

}
}
public class CommonN16 : IHotfixMessage
{
public short Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN16(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeN16(rBuffer, ref nOffset);

}
}
public class CommonN32 : IHotfixMessage
{
public int Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN32(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);

}
}
public class CommonN64 : IHotfixMessage
{
public long Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN64(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeN64(rBuffer, ref nOffset);

}
}
public class CommonU8 : IHotfixMessage
{
public byte Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeU8(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeU8(rBuffer, ref nOffset);

}
}
public class CommonU16 : IHotfixMessage
{
public ushort Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeU16(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeU16(rBuffer, ref nOffset);

}
}
public class CommonU32 : IHotfixMessage
{
public uint Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeU32(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeU32(rBuffer, ref nOffset);

}
}
public class CommonU64 : IHotfixMessage
{
public ulong Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeU64(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeU64(rBuffer, ref nOffset);

}
}
public class CommonF32 : IHotfixMessage
{
public float Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeF32(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeF32(rBuffer, ref nOffset);

}
}
public class CommonF64 : IHotfixMessage
{
public double Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeF64(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeF64(rBuffer, ref nOffset);

}
}
public class CommonStr : IHotfixMessage
{
public string Value;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeSTR(this.Value, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value = NetworkMessageEncode.DeserializeSTR(rBuffer, ref nOffset);

}
}
public class T1 : IHotfixMessage
{
public int Value1;
public float Value2;
public string Value3;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN32(this.Value1, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeF32(this.Value2, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeSTR(this.Value3, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value1 = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
this.Value2 = NetworkMessageEncode.DeserializeF32(rBuffer, ref nOffset);
this.Value3 = NetworkMessageEncode.DeserializeSTR(rBuffer, ref nOffset);

}
}
public class T2 : IHotfixMessage
{
public int Value1;
public float Value2;
public string Value3;
public int Value4;
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeN32(this.Value1, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeF32(this.Value2, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeSTR(this.Value3, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeN32(this.Value4, rBuffer, ref nOffset);

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value1 = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
this.Value2 = NetworkMessageEncode.DeserializeF32(rBuffer, ref nOffset);
this.Value3 = NetworkMessageEncode.DeserializeSTR(rBuffer, ref nOffset);
this.Value4 = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);

}
}
public class Test : IHotfixMessage
{
public bool Value0;
public byte Value1;
public short Value2;
public int Value3;
public long Value4;
public byte Value5;
public ushort Value6;
public uint Value7;
public ulong Value8;
public float Value9;
public double Value10;
public string Value11;
public T1 Value12 = new T1();
public List<T2> Value13 = new List<T2>();
public List<int> Value14 = new List<int>();
public List<string> Value15 = new List<string>();
public void Serialize(byte[] rBuffer, int nSize, ref int nOffset)
{
NetworkMessageEncode.SerializeBOOL(this.Value0, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeN8(this.Value1, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeN16(this.Value2, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeN32(this.Value3, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeN64(this.Value4, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeU8(this.Value5, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeU16(this.Value6, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeU32(this.Value7, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeU64(this.Value8, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeF32(this.Value9, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeF64(this.Value10, rBuffer, ref nOffset);
NetworkMessageEncode.SerializeSTR(this.Value11, rBuffer, ref nOffset);
this.Value12.Serialize(rBuffer, nSize - nOffset, ref nOffset);
int nValue13Count = this.Value13.Count;
NetworkMessageEncode.SerializeN32(nValue13Count, rBuffer, ref nOffset);
for (int i = 0; i < nValue13Count; i++){this.Value13[i].Serialize(rBuffer, nSize - nOffset, ref nOffset);}
int nValue14Count = this.Value14.Count;
NetworkMessageEncode.SerializeN32(nValue14Count, rBuffer, ref nOffset);
for (int i = 0; i < nValue14Count; i++){NetworkMessageEncode.SerializeN32(this.Value14[i], rBuffer, ref nOffset);}
int nValue15Count = this.Value15.Count;
NetworkMessageEncode.SerializeN32(nValue15Count, rBuffer, ref nOffset);
for (int i = 0; i < nValue15Count; i++){NetworkMessageEncode.SerializeSTR(this.Value15[i], rBuffer, ref nOffset);}

}
public void Deserialize(byte[] rBuffer, int nSize, ref int nOffset)
{
this.Value0 = NetworkMessageEncode.DeserializeBOOL(rBuffer, ref nOffset);
this.Value1 = NetworkMessageEncode.DeserializeN8(rBuffer, ref nOffset);
this.Value2 = NetworkMessageEncode.DeserializeN16(rBuffer, ref nOffset);
this.Value3 = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
this.Value4 = NetworkMessageEncode.DeserializeN64(rBuffer, ref nOffset);
this.Value5 = NetworkMessageEncode.DeserializeU8(rBuffer, ref nOffset);
this.Value6 = NetworkMessageEncode.DeserializeU16(rBuffer, ref nOffset);
this.Value7 = NetworkMessageEncode.DeserializeU32(rBuffer, ref nOffset);
this.Value8 = NetworkMessageEncode.DeserializeU64(rBuffer, ref nOffset);
this.Value9 = NetworkMessageEncode.DeserializeF32(rBuffer, ref nOffset);
this.Value10 = NetworkMessageEncode.DeserializeF64(rBuffer, ref nOffset);
this.Value11 = NetworkMessageEncode.DeserializeSTR(rBuffer, ref nOffset);
this.Value12.Deserialize(rBuffer, nSize - nOffset, ref nOffset);
int nValue13Count = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
for(int i = 0; i < nValue13Count; i++)
{
T2 temp = new T2();temp.Deserialize(rBuffer, nSize - nOffset, ref nOffset);this.Value13.Add(temp);
}
int nValue14Count = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
for(int i = 0; i < nValue14Count; i++)
{
this.Value14.Add(NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset));
}
int nValue15Count = NetworkMessageEncode.DeserializeN32(rBuffer, ref nOffset);
for(int i = 0; i < nValue15Count; i++)
{
this.Value15.Add(NetworkMessageEncode.DeserializeSTR(rBuffer, ref nOffset));
}

}
}
public class HotfixNetOpcode
{
public const int C2S_LOGIN = 1;
public const int S2C_LOGIN_RESP = 2;
}

}
